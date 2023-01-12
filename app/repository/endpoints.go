package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/Anton-Hudz/MovieList/app/entities"
	"github.com/Anton-Hudz/MovieList/app/globals"
	"github.com/lib/pq"
	// _ "github.com/lib/pq"
)

const ErrCodeUniqueViolation = "unique_violation"

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) AddUser(user entities.User) (int, error) {
	var id int
	SQL := fmt.Sprintf(`INSERT INTO %s (login, password_hash, age, user_role) values ($1, $2, $3, $4) RETURNING id`, usersTable)

	if err := r.DB.QueryRow(SQL, user.Login, user.Password, user.Age, user.User_Role).Scan(&id); err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return 0, globals.ErrDuplicateLogin
		}

		return 0, fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}

func (r *Repo) GetUser(login, password string) (entities.User, error) {
	var user entities.User
	SQL := fmt.Sprintf(`SELECT id, user_role FROM %s WHERE login=$1 AND password_hash=$2`, usersTable)

	if err := r.DB.QueryRow(SQL, login, password).Scan(&user.ID, &user.User_Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, globals.ErrNotFound
		}

		return entities.User{}, fmt.Errorf("internal error while scanning row: %w", err)
	}

	return user, nil
}

func (r *Repo) AddToken(userToken string, user entities.User) error {
	SQL := fmt.Sprintf(`UPDATE %s SET token = $1, deleted_token = null WHERE id = $2`, usersTable)

	if _, err := r.DB.Exec(SQL, userToken, user.ID); err != nil {
		errors.Is(err, sql.ErrNoRows)
		return globals.ErrNotFound
	}

	return nil
}

type CheckToken struct {
	Token string `json:"token"`
}

func (r *Repo) CheckToken(accessToken string) error {
	var user CheckToken
	SQL := fmt.Sprintf(`SELECT token FROM %s WHERE token=$1 AND deleted_token is null`, usersTable)

	if err := r.DB.QueryRow(SQL, accessToken).Scan(&user.Token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return globals.ErrTokenIsAlreadyDeleted
		}

		return fmt.Errorf("internal error while scanning row: %w", err)
	}

	return nil
}

func (r *Repo) DeleteToken(userId int, token string) error {
	SQL := fmt.Sprintf(`UPDATE %s SET deleted_token = NOW() WHERE id = $1 AND token = $2`, usersTable)

	if _, err := r.DB.Exec(SQL, userId, token); err != nil {
		errors.Is(err, sql.ErrNoRows)
		return globals.ErrTokenIsAlreadyDeleted
	}

	return nil
}

func (r *Repo) GetDirectorId(film entities.Film) (int, error) {
	var id int
	SQL := fmt.Sprintf(`SELECT id FROM %s WHERE name=$1`, directorTable)

	if err := r.DB.QueryRow(SQL, film.Director_Name).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, globals.ErrNotFound
		}

		return 0, fmt.Errorf("internal error while scanning row: %w", err)
	}

	return id, nil
}

func (r *Repo) AddMovie(film entities.Film, directorId int) (int, error) {
	var id int

	rateInt, err := strconv.ParseFloat(film.Rate, 32)
	if err != nil {
		return 0, errors.New("error rate must be number")
	}
	yearInt, err := strconv.Atoi(film.Year)
	if err != nil {
		return 0, errors.New("error year must be number")
	}
	minutesInt, err := strconv.Atoi(film.Minutes)
	if err != nil {
		return 0, errors.New("error duration must be number")
	}

	SQL := fmt.Sprintf(`INSERT INTO %s (name, genre, director_id, rate, year, minutes) values ($1, $2, $3, $4, $5, $6) RETURNING id`, filmTable)

	if err := r.DB.QueryRow(SQL, film.Name, film.Genre, directorId, rateInt, yearInt, minutesInt).Scan(&id); err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return 0, globals.ErrDuplicateFilmName
		}

		return 0, fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}
