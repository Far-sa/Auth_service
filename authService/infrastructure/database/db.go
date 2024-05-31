package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // Import Postgres driver (if using Postgres)
)

type SqlDB struct {
	db *sql.DB
}

// Add this method to the sqlDB type.
func (s *SqlDB) Conn() *sql.DB {
	return s.db
}

func NewSQLDB(url string) (*SqlDB, error) {

	// dsn := "postgres://root:password@postgres-auth:5432/auth-db?sslmode=disable" // Connect to the database directly
	db, err := sql.Open("postgres", url)
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

// Reusable functions for common database interactions using prepared statements
// func (db *SqlDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	stmt, err := db.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close() // Close prepared statement after execution

// 	result, err := stmt.ExecContext(ctx, args...)
// 	return result, err
// }

// func (db *SqlDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
// 	stmt, err := db.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil
// 	}
// 	defer stmt.Close() // Close prepared statement after execution

// 	return stmt.QueryRowContext(ctx, args...)
// }

// ... additional reusable functions for specific needs (e.g., scanning rows)
