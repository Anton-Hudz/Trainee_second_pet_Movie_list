package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/cfg"
	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

const (
	up           = "up"
	down         = "down"
	migrationDir = "./migration"
)

func main() {
	config, err := cfg.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %+v", err)

		return
	}

	db, err := repository.ConnectDB(config.DB)
	if err != nil {
		fmt.Printf("Error connecting to database on host: %s, port: %s, with error: %s", config.DB.Host, config.DB.Port, err)

		return
	}

	direction := flag.String("migrate", "", "applying migration direction")
	flag.Parse()

	if *direction != up && *direction != down {
		log.Printf("Wrong flag provided, choose '-migrate %s' or '-migrate %s'\n", up, down)

		return
	}

	if err := migrateDB(db, *direction); err != nil {
		fmt.Printf("Failed making migrations: %v", err)
	}
}

func migrateDB(db *sql.DB, direction string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationDir,
	}

	var dir migrate.MigrationDirection
	if direction == down {
		dir = 1
	}

	n, err := migrate.Exec(db, "postgres", migrations, dir)
	if err != nil {
		return fmt.Errorf("migration up failed: %w", err)
	}
	log.Printf("Number of applied migration is: %d", n)

	return nil
}