package usecase

import (
	// "fmt"

	"errors"
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

func (a *AuthUser) GenerateToken(login, password string) (string, error) {
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
	//передать токен в БД  и записать его
	// вызвать что то типа AddToken(token string, user.ID int)
	// а в репозитории формируем отбор по ид и записываем токен в ячейку (добавить в миграции ячейку
	//с админом, токеном, ячейку удаленный токен)
	// переименовать метод, он не только генерирует токен

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

	//возможно, нужно проверить, рабочий ли токен не выполнен ли по нему лог-аут
	// передав через ЮзКейс роспарсенный ИД в репозиторий и проверив его статус в БД
	// к примеру ячейка удаленный токен не пустая означает что токен не валидный
	//и есть еще вариант, сверять токен в каждой функции при попытке записать фильм или ТД

	return claims.UserID, nil
}
