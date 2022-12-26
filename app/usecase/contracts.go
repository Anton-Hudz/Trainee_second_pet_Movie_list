package usecase

import "github.com/Anton-Hudz/MovieList/app/repository"

type UserUseCase interface {
}

type FilmUseCase interface {
}

type UseCase struct {
	UserUseCase
	FilmUseCase
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{}
}
