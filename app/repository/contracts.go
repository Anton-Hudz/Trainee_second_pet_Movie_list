// package repository
package repository

type UserRepository interface {
}

type FilmRepository interface {
}

type Repository struct {
	UserRepository
	FilmRepository
}

func NewRepository() *Repository {
	return &Repository{}
}
