package usecase

import (
	// "fmt"

	"errors"
	"fmt"
	"time"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/pkg/hash"
	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL   = 30 * time.Minute
	signingKey = "gdajkl156alaflkj"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	UserRole string `json:"user_role"`
}

type AuthUser struct {
	Repo repository.UserRepository
}

func NewAuthUser(repo repository.UserRepository) *AuthUser {
	return &AuthUser{Repo: repo}
}

func (a *AuthUser) AddUser(user entities.User) (int, error) {
	user.Password = hash.GeneratePasswordHash(user.Password)

	return a.Repo.AddUser(user)
}

func (a *AuthUser) GenerateAddToken(login, password string) (string, error) {
	user, err := a.Repo.GetUser(login, hash.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.User_Role,
	})

	userToken, _ := token.SignedString([]byte(signingKey))

	err = a.Repo.AddToken(userToken, user)
	if err != nil {
		return "", fmt.Errorf("error occured while added token to database: %w", err)
	}

	return userToken, nil
}

func (a *AuthUser) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	if err := a.Repo.CheckToken(accessToken); err != nil {
		return 0, "", fmt.Errorf("error while checking token from database: %w", err)
	}

	return claims.UserID, claims.UserRole, nil
}

func (a *AuthUser) SignOut(userId int, token string) error {
	err := a.Repo.DeleteToken(userId, token)
	if err != nil {
		return fmt.Errorf("error while deleting token from database: %w", err)
	}

	return nil
}

type FilmService struct {
	Repo repository.FilmRepository
}

func NewFilmService(repo repository.FilmRepository) *FilmService {
	return &FilmService{Repo: repo}
}

func (f *FilmService) ValidateFilmData(film entities.Film) error {

	if film.Minutes <= 0 {
		return errors.New("error while checking film length, length must be above 0 minutes")
	}

	// if err := f.Repo.CheckUniqueFilm(film); err != nil {
	// 	return fmt.Errorf("error while checking uniqueness of film in database", err)
	// }

	return nil
}

func (f *FilmService) GetDirectorId(film entities.Film) (int, error) {
	id, err := f.Repo.GetDirectorId(film)
	if err != nil {
		return 0, fmt.Errorf("error occured while getting director ID: %w", err)
	}
	return id, nil
}

func (f *FilmService) AddFilm(film entities.Film, directorId int) (int, error) {
	id, err := f.Repo.AddMovie(film, directorId)
	if err != nil {
		return 0, fmt.Errorf("error while added movie to database: %w", err)
	}
	return id, nil

}
