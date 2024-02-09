package repository

import (
	"database/sql"
	"time"

	"github.com/qvtec/go-app/internal/domain"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetByID(id int) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id int) error
}

type mysqlUserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
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
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := r.DB.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	user.ID = int(userID)

	return nil
}

func (r *mysqlUserRepository) GetByID(id int) (*domain.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	user := &domain.User{}
	var createdAt, updatedAt, deletedAt sql.NullString
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
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

	user.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt.String)
	user.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt.String)
	if deletedAt.Valid {
		user.DeletedAt, _ = time.Parse("2006-01-02 15:04:05", deletedAt.String)
	}

	return user, nil
}

func (r *mysqlUserRepository) Update(user *domain.User) error {
	currentTime := time.Now().UTC()
	user.UpdatedAt = currentTime
	query := "UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?"
	result, err := r.DB.Exec(query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected != 1 {
		return domain.ErrRowsAffected
	}

	return nil
}

func (r *mysqlUserRepository) Delete(id int) error {
	currentTime := time.Now().UTC()
	query := "UPDATE users SET deleted_at = ? WHERE id = ?"
	result, err := r.DB.Exec(query, currentTime, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected != 1 {
		return domain.ErrRowsAffected
	}

	return nil
}
