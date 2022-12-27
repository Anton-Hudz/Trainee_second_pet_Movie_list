// package repository
package repository

import "database/sql"

type UserRepository interface {
}

type FilmRepository interface {
}

type Repository struct {
	UserRepository
	FilmRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
