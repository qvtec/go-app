package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrAlbumNotFound = errors.New("album not found")
)