package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrRowsAffected = errors.New("unexpected number of rows affected")
)