package globals

import "errors"

var (
	ErrDuplicateLogin        = errors.New("login is already taken")
	ErrDuplicateFilmName     = errors.New("film is already taken")
	ErrNotFound              = errors.New("no data found in DB")
	ErrTokenIsAlreadyDeleted = errors.New("token is already deleted")
)
