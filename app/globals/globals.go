package globals

import "errors"

var (
	ErrDuplicateLogin        = errors.New("login is already taken")
	ErrDuplicateFilmName     = errors.New("film is already taken")
	ErrNotFound              = errors.New("no data found in DB")
	ErrTokenIsAlreadyDeleted = errors.New("token is already deleted")
	ErrWrongMovieName        = errors.New("wrong movie name or the movie is not in the database")
	ErrDuplicateMovieInList  = errors.New("duplicate movie in list")
	ErrIncorrectUserData     = errors.New("error login must be phone number in format: 0671234567")
)
