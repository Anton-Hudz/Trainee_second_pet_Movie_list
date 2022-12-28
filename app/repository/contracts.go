// package repository
package repository

import (
	"database/sql"

	"github.com/Anton-Hudz/MovieList/app/entities"
)

type UserRepository interface {
	AddUser(user entities.User) (int, error)
}

type FilmRepository interface {
}

type Repository struct {
	UserRepository
	FilmRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewRepo(db),
	}
}
