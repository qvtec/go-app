package domain

import "time"

type Album struct {
	ID         int
	Title      string
	Contents   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
