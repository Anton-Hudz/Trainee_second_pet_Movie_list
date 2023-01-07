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
	tokenTTL   = 5 * time.Minute
	signingKey = "gdajkl156alaflkj"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
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
	})

	userToken, _ := token.SignedString([]byte(signingKey))

	err = a.Repo.AddToken(userToken, user)
	if err != nil {
		return "", fmt.Errorf("error occured while added token to database: %w", err)
	}

	return userToken, nil
}

func (a *AuthUser) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	//возможно(ОБЯЗАТЕЛЬНО), нужно проверить, рабочий ли токен не выполнен ли по нему лог-аут
	// передав через ЮзКейс роспарсенный ИД в репозиторий и проверив его статус в БД
	// к примеру ячейка удаленный токен не пустая означает что токен не валидный
	//и есть еще вариант, сверять токен в каждой функции при попытке записать фильм или ТД

	return claims.UserID, nil
}

func (a *AuthUser) SignOut(userId int, token string) error {
	err := a.Repo.DeleteToken(userId, token)
	if err != nil {
		return fmt.Errorf("error while deleting token from database: %w", err)
	}

	return nil
}
