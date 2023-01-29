package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/cfg"
	"github.com/Anton-Hudz/MovieList/logger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

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

	if err := logger.CustomLogger(config.LogLevel); err != nil {
		fmt.Printf("Failed to create logger: %s", err)

		return
	}

	db, err := repository.ConnectDB(config.DB)
	if err != nil {
		logrus.Errorf("Error connecting to database on host: %s, port: %s, with error: %s", config.DB.Host, config.DB.Port, err)

		return
	}

	direction := flag.String("migrate", "", "applying migration direction")
	flag.Parse()

	if *direction != up && *direction != down {
		logrus.Errorf("Wrong flag provided, choose '-migrate %s' or '-migrate %s'\n", up, down)

		return
	}

	if err := migrateDB(db, *direction); err != nil {
		logrus.Errorf("Failed making migrations: %v", err)

		return
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
	logrus.Infof("Number of applied migration is: %d", n)

	return nil
}
