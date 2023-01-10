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
}

type UseCase struct {
	UserUseCase
	FilmUseCase
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		UserUseCase: NewAuthUser(repos.UserRepository),
	}
}
