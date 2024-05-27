package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type SqlDB struct {
	db *sql.DB
}

// Add this method to the sqlDB type.
func (d *SqlDB) Conn() *sql.DB {
	return d.db
}

func NewSQLDB(pathUrl string) (*SqlDB, error) {

	//dsn := "postgres://root:password@postgres-auth:5432/auth-db?sslmode=disable" // Connect to the database directly
	db, err := sql.Open("postgres", pathUrl)
	if err != nil {
		panic(fmt.Errorf("can not open postgres database: %v", err))
	}

	// Create the database if it doesn't exist.
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS auth-db")
	if err != nil {
		// If the error is due to the database already existing, ignore it.
		if !strings.Contains(err.Error(), "already exists") {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
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
