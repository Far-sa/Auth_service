package database

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "created_at", "birthdate"}).
		AddRow("1", "John Doe", "johndoe", "john@example.com", time.Now(), time.Now())

	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
		WithArgs("1").
		WillReturnRows(rows)

	sqlDB := &SqlDB{db: db}

	userProfile, err := sqlDB.GetUserByID(context.Background(), "1")

	assert.NoError(t, err)
	assert.NotNil(t, userProfile)
	assert.Equal(t, "John Doe", userProfile.FullName)
	assert.Equal(t, "johndoe", userProfile.Username)
	assert.Equal(t, "john@example.com", userProfile.Email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByIDScanError(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
		WithArgs("invalid").
		WillReturnError(sqlmock.ErrConnDone)

	sqlDB := &SqlDB{db: db}

	userProfile, err := sqlDB.GetUserByID(context.Background(), "invalid")

	assert.Error(t, err)
	assert.Nil(t, userProfile)
	assert.Contains(t, err.Error(), "error scanning user")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByIDPostgresError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pgErr := &pq.Error{
		Code: "23505",
	}

	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
		WithArgs("conflict").
		WillReturnError(pgErr)

	sqlDB := &SqlDB{db: db}

	userProfile, err := sqlDB.GetUserByID(context.Background(), "conflict")

	assert.Error(t, err)
	assert.Nil(t, userProfile)
	assert.Contains(t, err.Error(), "error getting user")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
