package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type SqlDB struct {
	db *sql.DB
}

// Add this method to the sqlDB type.
func (d *SqlDB) Conn() *sql.DB {
	return d.db
}

func NewSQLDB(dsn string) (*SqlDB, error) {

	//dsn := "postgres://root:password@postgres-auth:5432/auth-db?sslmode=disable" // Connect to the database directly
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("can not open postgres database: %v", err))
	}

	// Set connection pool parameters (optional)
	db.SetMaxOpenConns(20) // Adjust as needed
	db.SetMaxIdleConns(10) // Adjust as needed

	// Perform initial ping to check connectivity
	if err := db.PingContext(context.Background()); err != nil {
		return nil, errors.New("failed to ping database: " + err.Error())
	}

	return &SqlDB{db: db}, nil
}
