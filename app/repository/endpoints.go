package repository

import (
	"database/sql"
	"fmt"

	"github.com/Anton-Hudz/MovieList/app/entities"
	_ "github.com/lib/pq"
)

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
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, age) values ($1, $2, $3) RETURNING id", userTable)
	row := r.DB.QueryRow(query, user.Login, user.Password, user.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// INSERT INTO film (name, genre, director_id, rate, year, minutes)
// 			VALUES ('Avatar', 'Fantastic', 2, 9, 2009, 155)
