package db_test

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qvtec/go-app/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")

	os.Setenv("DATABASE_NAME", "testing")
	m.Run()

	os.Setenv("DATABASE_HOST", host)
	os.Setenv("DATABASE_PORT", port)
	os.Setenv("DATABASE_NAME", dbName)
	os.Setenv("DATABASE_USER", user)
	os.Setenv("DATABASE_PASSWORD", password)
}

func TestNewMySQLDB(t *testing.T) {
	password := os.Getenv("DATABASE_PASSWORD")

	t.Run("OK", func(t *testing.T) {
		db, err := db.NewMySQLDB()
		assert.NoError(t, err)
		defer db.Close()
	})

	t.Run("Open is nil InvalidEnv", func(t *testing.T) {
		os.Setenv("DATABASE_PASSWORD", "invalid_password")
		db, err := db.NewMySQLDB()
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("Open is nil Unsetenv", func(t *testing.T) {
		os.Unsetenv("DATABASE_PASSWORD")
		db, err := db.NewMySQLDB()
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	os.Setenv("DATABASE_PASSWORD", password)
}
