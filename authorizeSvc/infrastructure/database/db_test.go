package database_test

import (
	"authorization-service/infrastructure/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewSQLDB(t *testing.T) {
	// Create a new SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %s", err)
	}
	defer db.Close()

	// Mock the PingContext method
	mock.ExpectPing().WillReturnError(nil)

	// Call the NewSQLDB function
	sqlDB, err := database.NewSQLDB("mock DSN")

	// Assert that there was no error and that the SqlDB object was created
	assert.NoError(t, err)
	assert.NotNil(t, sqlDB)

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
