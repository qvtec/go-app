package repository

import (
	"database/sql"
	"time"

	"github.com/qvtec/go-app/internal/domain"
	"github.com/qvtec/go-app/pkg/db"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetByID(id int) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id int) error
}

type mysqlUserRepository struct {
	DB *db.MySQLDB
}

func NewUserRepository(db *db.MySQLDB) UserRepository {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (r *mysqlUserRepository) GetAll() ([]*domain.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *mysqlUserRepository) Create(user *domain.User) error {
	currentTime := time.Now().UTC()
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	query := "INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Execute(query, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(userID)

	return nil
}

func (r *mysqlUserRepository) GetByID(id int) (*domain.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *mysqlUserRepository) Update(user *domain.User) error {
	currentTime := time.Now().UTC()
	user.UpdatedAt = currentTime
	query := "UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?"
	_, err := r.DB.Execute(query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Execute(query, id)
	if err != nil {
		return err
	}

	return nil
}
