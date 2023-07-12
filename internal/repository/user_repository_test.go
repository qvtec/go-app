package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qvtec/go-app/internal/domain"
	"github.com/qvtec/go-app/internal/repository"
	"github.com/qvtec/go-app/pkg/db"
)

func TestMySQLUserRepository(t *testing.T) {
	db, err := db.NewMySQLDB()
	assert.NoError(t, err)
	defer db.Close()

	tx, err := db.Begin()
	assert.NoError(t, err)
	defer tx.Rollback()

	userRepo := repository.NewUserRepository(db)

	user := &domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Create
	err = userRepo.Create(user)
	assert.NoError(t, err)

	// GetByID (Create確認)
	fetchedUser, err := userRepo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Name, fetchedUser.Name)
	assert.Equal(t, user.Email, fetchedUser.Email)

	// Update
	user.Name = "Jane Doe"
	err = userRepo.Update(user)
	assert.NoError(t, err)

	// GetByID (UPDATE確認)
	fetchedUser, err = userRepo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Name, fetchedUser.Name)
	assert.Equal(t, user.Email, fetchedUser.Email)

	// GetAll
	users, err := userRepo.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, users)

	// Delete
	err = userRepo.Delete(user.ID)
	assert.NoError(t, err)

	// GetByID (DELETE確認)
	fetchedUser, err = userRepo.GetByID(user.ID)
	assert.Error(t, err)
	assert.Nil(t, fetchedUser)
}
