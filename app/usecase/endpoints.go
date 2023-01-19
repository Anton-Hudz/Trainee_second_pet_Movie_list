package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/pkg/hash"
	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL   = 12 * time.Hour
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
	Repo repository.FilmRepository
}

func NewFilmService(repo repository.FilmRepository) *FilmService {
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

func (f *FilmService) GetFilmID(filmName string) (int, error) {
	id, err := f.Repo.GetFilmID(filmName)
	if err != nil {
		return 0, fmt.Errorf("error occured while getting film ID: %w", err)
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

func (f *FilmService) AddFilmToFavourite(userID any, filmID int) (int, error) {
	id, err := f.Repo.AddMovieToList(userID, filmID, repository.FavouriteTable)
	if err != nil {
		return 0, fmt.Errorf("error occured while added movie to favourite list: %w", err)
	}

	return id, nil
}

func (f *FilmService) AddToWishlist(userID any, filmID int) (int, error) {
	id, err := f.Repo.AddMovieToList(userID, filmID, repository.WishlistTable)
	if err != nil {
		return 0, fmt.Errorf("error occured while added movie to wish list: %w", err)
	}

	return id, nil
}
