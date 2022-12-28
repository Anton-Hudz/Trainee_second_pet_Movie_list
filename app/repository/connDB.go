package repository

import (
	"database/sql"
	"fmt"

	"github.com/Anton-Hudz/MovieList/cfg"
)

const (
	userTable      = "users"
	filmTable      = "film"
	directorTable  = "director"
	favouriteTable = "favourite"
	wishlistTable  = "wishlist"
)

func ConnectDB(cfg cfg.DB) (*sql.DB, error) {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection:%w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %w", err)
	}

	return db, nil
}
