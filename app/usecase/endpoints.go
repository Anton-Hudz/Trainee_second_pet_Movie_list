package usecase

import (
	"fmt"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/pkg/hash"
)

type AuthUser struct {
	Repo repository.UserRepository
}

func NewAuthUser(repo repository.UserRepository) *AuthUser {
	return &AuthUser{Repo: repo}
}

func (a *AuthUser) AddUser(user entities.User) (int, error) {
	var err error
	user.Password, err = hash.GeneratePasswordHash(user.Password)
	if err != nil {
		return 0, fmt.Errorf("error while adding user to database: %w", err)
	}

	return a.Repo.AddUser(user)
}
