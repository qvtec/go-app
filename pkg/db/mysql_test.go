package db_test

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qvtec/go-app/pkg/db"
	"github.com/stretchr/testify/assert"
)

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
