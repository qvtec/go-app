package domain

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
