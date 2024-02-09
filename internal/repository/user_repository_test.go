package repository_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/qvtec/go-app/internal/domain"
	"github.com/qvtec/go-app/internal/repository"
)

var Users = []*domain.User{
	{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	},
	{
		ID:    2,
		Name:  "Jane Smith",
		Email: "jane@example.com",
	},
	{
		ID:    3,
		Name:  "Alice Johnson",
		Email: "alice@example.com",
	},
}

func setupDB() (*sql.DB, repository.UserRepository, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := "testing"

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, nil, err
	}

	userRepo := repository.NewUserRepository(db)

	return db, userRepo, nil
}

func TestUserRepository_Create(t *testing.T) {
	db, userRepo, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("OK", func(t *testing.T) {
		user := Users[0]
		err = userRepo.Create(user)
		assert.NoError(t, err)
	})

	t.Run("NG", func(t *testing.T) {
		user := Users[0]
		err = userRepo.Create(user)
		assert.Error(t, err)
	})
}

func TestUserRepository_GetAll(t *testing.T) {
	db, userRepo, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()

	user1 := Users[0]

	user2 := Users[1]
	err = userRepo.Create(user2)
	assert.NoError(t, err)

	users, err := userRepo.GetAll()
	assert.NoError(t, err)

	fetchedUser1, err := userRepo.GetByID(user1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser1)
	assert.Equal(t, users[0].ID, fetchedUser1.ID)
	assert.Equal(t, users[0].Name, fetchedUser1.Name)
	assert.Equal(t, users[0].Email, fetchedUser1.Email)

	fetchedUser2, err := userRepo.GetByID(user2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser2)
	assert.Equal(t, users[1].ID, fetchedUser2.ID)
	assert.Equal(t, users[1].Name, fetchedUser2.Name)
	assert.Equal(t, users[1].Email, fetchedUser2.Email)
}

func TestUserRepository_GetByID(t *testing.T) {
	db, userRepo, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()

	user := Users[0]

	t.Run("OK", func(t *testing.T) {
		fetchedUser, err := userRepo.GetByID(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedUser)
		assert.Equal(t, user.ID, fetchedUser.ID)
		assert.Equal(t, user.Name, fetchedUser.Name)
		assert.Equal(t, user.Email, fetchedUser.Email)
	})

	t.Run("ErrUserNotFound", func(t *testing.T) {
		_, err := userRepo.GetByID(100)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrUserNotFound, err)
	})

	t.Run("ErrUserNotFound", func(t *testing.T) {
		_, err := userRepo.GetByID(100)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrUserNotFound, err)
	})
}

func TestUserRepository_Update(t *testing.T) {
	db, userRepo, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("OK", func(t *testing.T) {
		user := Users[0]

		user.Name = "Jane Doe"
		err = userRepo.Update(user)
		assert.NoError(t, err)

		fetchedUser, err := userRepo.GetByID(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedUser)
		assert.Equal(t, user.ID, fetchedUser.ID)
		assert.Equal(t, user.Name, fetchedUser.Name)
		assert.Equal(t, user.Email, fetchedUser.Email)
		assert.NotEqual(t, fetchedUser.CreatedAt, fetchedUser.UpdatedAt)
	})

	t.Run("NG", func(t *testing.T) {
		user := Users[0]
		user.ID = 100
		err = userRepo.Update(user)
		assert.Error(t, err)
	})
}

func TestUserRepository_Delete(t *testing.T) {
	db, userRepo, err := setupDB()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("OK", func(t *testing.T) {
		user := Users[2]
		err = userRepo.Create(user)
		assert.NoError(t, err)

		err = userRepo.Delete(user.ID)
		assert.NoError(t, err)

		fetchedUser, err := userRepo.GetByID(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedUser.DeletedAt)
	})

	t.Run("NG", func(t *testing.T) {
		err = userRepo.Delete(100)
		assert.Error(t, err)
	})
}

//////////////////////////////////////
// DB Connect NG
//////////////////////////////////////

func setupDBNG() (*sql.DB, repository.UserRepository, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := "invalid_password"
	dbName := "testing"

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, nil, err
	}

	userRepo := repository.NewUserRepository(db)

	return db, userRepo, nil
}

func TestUserRepository_GetAll_InvalidDB(t *testing.T) {
	db, userRepo, err := setupDBNG()
	assert.NoError(t, err)
	defer db.Close()

	_, err = userRepo.GetAll()
	assert.Error(t, err)
}

func TestUserRepository_Create_InvalidDB(t *testing.T) {
	db, userRepo, err := setupDBNG()
	assert.NoError(t, err)
	defer db.Close()

	user := Users[0]
	err = userRepo.Create(user)
	assert.Error(t, err)
}

func TestUserRepository_Update_InvalidDB(t *testing.T) {
	db, userRepo, err := setupDBNG()
	assert.NoError(t, err)
	defer db.Close()

	user := Users[0]
	err = userRepo.Update(user)
	assert.Error(t, err)
}

func TestUserRepository_Delete_InvalidDB(t *testing.T) {
	db, userRepo, err := setupDBNG()
	assert.NoError(t, err)
	defer db.Close()

	err = userRepo.Delete(1)
	assert.Error(t, err)
}
