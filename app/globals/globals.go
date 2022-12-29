package globals

import "errors"

var (
	ErrDuplicateLogin = errors.New("login is already taken")
	ErrNotFound       = errors.New("no user found in DB")
)
