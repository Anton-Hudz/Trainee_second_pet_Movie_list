// package repository
package repository

import (
	"database/sql"

	"github.com/Anton-Hudz/MovieList/app/entities"
)

type UserRepository interface {
	AddUser(user entities.User) (int, error)
	GetUser(login, password string) (entities.User, error)
	AddToken(userToken string, user entities.User) error
	CheckToken(accessToken string) error
	DeleteToken(userId int, token string) error
}

type FilmRepository interface {
	// CheckUniqueFilm(film entities.Film) error
	GetDirectorId(film entities.Film) (int, error)
	AddMovie(film entities.Film, directorId int) (int, error)
}

type Repository struct {
	UserRepository
	FilmRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewRepo(db),
		FilmRepository: NewRepo(db),
	}
}
