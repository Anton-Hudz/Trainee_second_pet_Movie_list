package usecase

import (
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
	AddMovieToList(userID any, filmName string, table string) (int, error)
	GetFilmById(id int) (entities.FilmResponse, error)
	GetAllFilms(SQL string) ([]entities.FilmResponse, error)
}
