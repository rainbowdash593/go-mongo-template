package exceptions

import (
	"errors"
)

var (
	ErrUnhandled = errors.New("unhandled exceptions")

	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)
