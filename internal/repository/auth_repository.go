package repository

import (
	"database/sql"
	"time"

	"github.com/qvtec/go-app/internal/domain"
	"github.com/qvtec/go-app/pkg/crypto"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	GetPasswordByEmail(email string) (string, error)
	UpdatePassword(email, password string) error
}

type MySQLAuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &MySQLAuthRepository{DB: db}
}

func (r *MySQLAuthRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := "SELECT * FROM users WHERE email = ? AND deleted_at is NULL"
	row := r.DB.QueryRow(query, email)

	user := &domain.User{}
	var createdAt, updatedAt, deletedAt sql.NullString
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *MySQLAuthRepository) GetPasswordByEmail(email string) (string, error) {
	query := "SELECT password FROM users WHERE email = ? AND deleted_at is NULL"
	row := r.DB.QueryRow(query, email)

	user := &domain.User{}
	err := row.Scan(
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", domain.ErrUserNotFound
		}
		return "", err
	}

	return user.Password, nil
}

func (r *MySQLAuthRepository) UpdatePassword(email, newPassword string) error {
	currentTime := time.Now().UTC()
	hashedPassword, err := crypto.HashPassword(newPassword)

	query := "UPDATE users SET password = ?, updated_at = ? WHERE email = ?"
	result, err := r.DB.Exec(query, hashedPassword, currentTime, email)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected != 1 {
		return domain.ErrRowsAffected
	}

	return nil
}
