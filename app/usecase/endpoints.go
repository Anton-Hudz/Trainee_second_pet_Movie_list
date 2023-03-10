package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/globals"
	"github.com/Anton-Hudz/MovieList/pkg/hash"
	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL       = 12 * time.Hour
	firstLoginItem = 43
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	UserRole string `json:"user_role"`
}

type AuthUser struct {
	Repo UserRepository
}

func NewAuthUser(repo UserRepository) *AuthUser {
	return &AuthUser{Repo: repo}
}

func (a *AuthUser) AddUser(user entities.User, salt string) (int, error) {
	user.Password = hash.GeneratePasswordHash(user.Password, salt)
	if err := validateUser(user); err != nil {
		return 0, fmt.Errorf("error occured while validate user data: %w", err)
	}
	return a.Repo.AddUser(user)
}
func validateUser(user entities.User) error {
	if user.Login[0] != firstLoginItem {
		return globals.ErrIncorrectUserData
	}
	_, err := strconv.Atoi(user.Login)
	if err != nil {
		return globals.ErrIncorrectUserData
	}

	return nil
}

func (a *AuthUser) GenerateAddToken(login, password, signingKey, salt string) (string, int, error) {
	user, err := a.Repo.GetUser(login, hash.GeneratePasswordHash(password, salt))
	if err != nil {
		return "", 0, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.UserRole,
	})

	userToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", 0, err
	}

	err = a.Repo.AddToken(userToken, user)
	if err != nil {
		return "", 0, fmt.Errorf("error occured while added token to database: %w", err)
	}

	return userToken, user.ID, nil
}

func (a *AuthUser) ParseToken(accessToken, signingKey string) (int, string, error) {
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
		return 0, "", fmt.Errorf("error occured while checking token from database: %w", err)
	}

	return claims.UserID, claims.UserRole, nil
}

func (a *AuthUser) SignOut(userId int, token string) error {
	err := a.Repo.DeleteToken(userId, token)
	if err != nil {
		return fmt.Errorf("error occured while deleting token from database: %w", err)
	}

	return nil
}

type FilmService struct {
	Repo FilmRepository
}

func NewFilmService(repo FilmRepository) *FilmService {
	return &FilmService{Repo: repo}
}

func (f *FilmService) ValidateFilmData(film entities.Film) error {
	minutesInt, err := strconv.Atoi(film.Minutes)
	if err != nil {
		return errors.New("error duration must be number")
	}

	if err := checkRequiredFields(film); err != nil {
		return err
	}

	if err := checkMovieDuration(minutesInt); err != nil {
		return err
	}

	return nil
}

func checkRequiredFields(film entities.Film) error {
	const minLength = 1
	if len(film.Name) < minLength {
		return fmt.Errorf("movie name field must not be empty")
	}

	if len(film.Genre) < minLength {
		return fmt.Errorf("genre field must not be empty")
	}

	if len(film.Rate) < minLength {
		return fmt.Errorf("rate field must not be empty")
	}

	if len(film.Year) < minLength {
		return fmt.Errorf("year field must not be empty")
	}

	if len(film.Minutes) < minLength {
		return fmt.Errorf("minutes field must not be empty")
	}

	return nil
}

func checkMovieDuration(filmDuration int) error {
	const minFilmDuration int = 0

	if filmDuration <= minFilmDuration {
		return fmt.Errorf("error occured while checking film length, length must be above %v minutes", minFilmDuration)
	}

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
		return 0, fmt.Errorf("error occured while added movie to database: %w", err)
	}

	return id, nil
}

func (f *FilmService) MakeQuery(params entities.QueryParams) (string, error) {
	SQL, err := createQuery(params)
	if err != nil {
		return "", fmt.Errorf("error occured while getting query parameters: %w", err)
	}

	return SQL, nil
}

func (f *FilmService) GetFilmList(SQL string) ([]entities.FilmResponse, error) {

	filmlist, err := f.Repo.GetAllFilms(SQL)
	if err != nil {
		return []entities.FilmResponse{}, fmt.Errorf("error with query parameters: %w", err)
	}

	return filmlist, nil
}

func (f *FilmService) GetFilmById(id int) (entities.FilmResponse, error) {
	film, err := f.Repo.GetFilmById(id)
	if err != nil {
		return entities.FilmResponse{}, fmt.Errorf("error occured while getting film from database: %w", err)
	}

	return film, nil
}

func (f *FilmService) AddFilmToList(userID any, filmName, table string) (int, error) {
	id, err := f.Repo.AddMovieToList(userID, filmName, table)
	if err != nil {
		return 0, fmt.Errorf("error occured while added movie to %v: %w", table, err)
	}

	return id, nil
}
