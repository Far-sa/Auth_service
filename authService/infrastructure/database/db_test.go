package database_test

import (
	"authentication-service/infrastructure/database"
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// // TestNewSQLDB tests the NewSQLDB function for successful connection.
// func TestNewSQLDB(t *testing.T) {
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer mockDB.Close()

// 	mock.ExpectPing()

// 	dsn := ""
// 	sdb, err := database.NewSQLDB(dsn)
// 	if sdb == nil || err != nil {
// 		t.Errorf("Expected non-nil *sqlDB and no error, got %v and %v", sdb, err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

func TestNewSQLDB_R(t *testing.T) {
	dsn := "postgres://root:password@postgres-auth:5432/authz-db?sslmode=disable"
	sdb, err := database.NewSQLDB(dsn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sdb == nil {
		t.Error("expected non-nil *SqlDB, got nil")
	}

	conn := sdb.Conn()
	if conn == nil {
		t.Error("expected non-nil *sql.DB, got nil")
	}

	if err := conn.PingContext(context.Background()); err != nil {
		t.Errorf("failed to ping database: %v", err)
	}
}
func TestNewSQLDB(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	mock.ExpectPing()

	dsn := ""
	sdb, err := database.NewSQLDB(dsn)
	if sdb == nil || err != nil {
		t.Errorf("Expected non-nil *SqlDB and no error, got %v and %v", sdb, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
