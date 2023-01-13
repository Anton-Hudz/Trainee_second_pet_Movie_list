package usecase

import (
	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/repository"
)

type UserUseCase interface {
	AddUser(user entities.User) (int, error)
	GenerateAddToken(login, password string) (string, error)
	ParseToken(token string) (int, string, error)
	SignOut(userId int, token string) error
}

type FilmUseCase interface {
	ValidateFilmData(film entities.Film) error
	GetDirectorId(film entities.Film) (int, error)
	AddFilm(film entities.Film, directorId int) (int, error)
	GetFilmID(filmName string) (int, error)
	AddFilmToFavourite(userID any, filmID int) (int, error)
	AddToWishlist(userID any, filmID int) (int, error)
}

type UseCase struct {
	UserUseCase
	FilmUseCase
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		UserUseCase: NewAuthUser(repos.UserRepository),
		FilmUseCase: NewFilmService(repos.FilmRepository),
	}
}
