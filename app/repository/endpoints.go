package repository

import (
	"database/sql"
	"errors"
	"fmt"

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
	SQL := fmt.Sprintf("INSERT INTO %s (login, password_hash, age, admin) values ($1, $2, $3, $4) RETURNING id", usersTable)

	if err := r.DB.QueryRow(SQL, user.Login, user.Password, user.Age, user.Admin).Scan(&id); err != nil {
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
	SQL := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", usersTable)

	if err := r.DB.QueryRow(SQL, login, password).Scan(&user.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, globals.ErrNotFound
		}

		return entities.User{}, fmt.Errorf("internal error while scanning row: %w", err)
	}

	return user, nil

}

func (r *Repo) AddToken(userToken string, user entities.User) error {
	SQL := fmt.Sprintf(`UPDATE %s SET token = $1 WHERE id = $2`, usersTable)

	if _, err := r.DB.Exec(SQL, userToken, user.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return globals.ErrNotFound
		}

		return fmt.Errorf("internal error while scanning row: %w", err)
	}

	return nil
}

// INSERT INTO film (name, genre, director_id, rate, year, minutes)
// 			VALUES ('Avatar', 'Fantastic', 2, 9, 2009, 155)
