package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	DB *sql.DB
}

func NewMySQLDB() (*MySQLDB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/docker")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MySQLDB{
		DB: db,
	}, nil
}

func (m *MySQLDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.DB.QueryRow(query, args...)
}

func (m *MySQLDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.Query(query, args...)
}

func (m *MySQLDB) Execute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}