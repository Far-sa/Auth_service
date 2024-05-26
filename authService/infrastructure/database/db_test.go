package database_test

import (
	"authentication-service/infrastructure/database"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// TestNewSQLDB tests the NewSQLDB function for successful connection.
func TestNewSQLDB(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	mock.ExpectPing()

	sdb, err := database.NewSQLDB()
	if sdb == nil || err != nil {
		t.Errorf("Expected non-nil *sqlDB and no error, got %v and %v", sdb, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
