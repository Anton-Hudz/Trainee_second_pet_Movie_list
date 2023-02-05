package usecase

import (
	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/repository"
)

type UserUseCase interface {
	AddUser(user entities.User, salt string) (int, error)
	GenerateAddToken(login, password, signingKey, salt string) (string, int, error)
	ParseToken(token, signingKey string) (int, string, error)
	SignOut(userId int, token string) error
}

type FilmUseCase interface {
	ValidateFilmData(film entities.Film) error
	GetDirectorId(film entities.Film) (int, error)
	AddFilm(film entities.Film, directorId int) (int, error)
	AddFilmToList(userID any, filmName, table string) (int, error)
	GetFilmById(id int) (entities.FilmResponse, error)
	MakeQuery(params entities.QueryParams) (string, error)
	GetFilmList(query string) ([]entities.FilmResponse, error)
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
