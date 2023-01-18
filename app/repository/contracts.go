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
	GetDirectorId(film entities.Film) (int, error)
	AddMovie(film entities.Film, directorId int) (int, error)
	GetFilmID(filmName string) (int, error)
	AddMovieToList(userID any, filmID int, table string) (int, error)
	GetFilmById(id int) (entities.FilmFromDB, error)
	GetAllFilms(SQL string) ([]entities.FilmFromDB, error)
	GetDirectorName(id int) (string, error)
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
