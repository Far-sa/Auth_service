package repository

// import (
// 	"context"
// 	"testing"
// 	"time"
// 	"user-service/infrastructure/database"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/lib/pq"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetUserByIDSuccess(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "created_at", "birthdate"}).
// 		AddRow("1", "John Doe", "johndoe", "john@example.com", time.Now(), time.Now())

// 	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
// 		WithArgs("1").
// 		WillReturnRows(rows)

// 	sqlDB := &database.SqlDB{db: db}

// 	userProfile, err := sqlDB.GetUserByID(context.Background(), "1")

// 	assert.NoError(t, err)
// 	assert.NotNil(t, userProfile)
// 	assert.Equal(t, "John Doe", userProfile.FullName)
// 	assert.Equal(t, "johndoe", userProfile.Username)
// 	assert.Equal(t, "john@example.com", userProfile.Email)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetUserByIDScanError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
// 		WithArgs("invalid").
// 		WillReturnError(sqlmock.ErrConnDone)

// 	sqlDB := &SqlDB{db: db}

// 	userProfile, err := sqlDB.GetUserByID(context.Background(), "invalid")

// 	assert.Error(t, err)
// 	assert.Nil(t, userProfile)
// 	assert.Contains(t, err.Error(), "error scanning user")

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetUserByIDPostgresError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	pgErr := &pq.Error{
// 		Code: "23505",
// 	}

// 	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
// 		WithArgs("conflict").
// 		WillReturnError(pgErr)

// 	sqlDB := &SqlDB{db: db}

// 	userProfile, err := sqlDB.GetUserByID(context.Background(), "conflict")

// 	assert.Error(t, err)
// 	assert.Nil(t, userProfile)
// 	assert.Contains(t, err.Error(), "error getting user")

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"
// 	"user-service/infrastructure/database"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetUserByID(t *testing.T) {
// 	// Create a new instance of sqlmock
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// Mock expected behavior
// 	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "birthdate", "created_at"}).
// 		AddRow("1", "John Doe", "johndoe", "john@example.com", time.Now(), time.Now())

// 	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
// 		WithArgs("1").
// 		WillReturnRows(rows)

// 	r := &DB{conn: &database.SqlDB{db: db}}

// 	// Call the method we want to test
// 	userProfile, err := r.GetUserByID(context.Background(), "1")

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.NotNil(t, userProfile)
// 	assert.Equal(t, "John Doe", userProfile.FullName)
// 	assert.Equal(t, "johndoe", userProfile.Username)
// 	assert.Equal(t, "john@example.com", userProfile.Email)

// 	// Ensure all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetUserByIDError(t *testing.T) {
// 	// Create a new instance of sqlmock
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// Mock expected behavior for an error case
// 	mock.ExpectQuery(`SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id = \$1`).
// 		WithArgs("nonexistent").
// 		WillReturnError(errors.New("user not found"))

// 	r := &DB{conn: db}

// 	// Call the method we want to test
// 	userProfile, err := r.GetUserByID(context.Background(), "nonexistent")

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, userProfile)
// 	assert.EqualError(t, err, "user not found")

// 	// Ensure all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetUserByID(t *testing.T) {
// 	// Create a mock DB connection
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer mockDB.Close()

// 	// Create a new instance of the repository with the mock DB connection
// 	repo := repository.DB{conn: mockDB}

// 	// Define the expected query and result
// 	expectedUserID := "123"
// 	expectedUser := &entity.UserProfile{
// 		ID:        expectedUserID,
// 		FullName:  "John Doe",
// 		Username:  "johndoe",
// 		Email:     "johndoe@example.com",
// 		Birthdate: time.Now(),
// 		CreatedAt: time.Now(),
// 	}
// 	mock.ExpectQuery("SELECT id, full_name, username, email, created_at, birthdate FROM users WHERE user_id =").
// 		WithArgs(expectedUserID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "full_name", "username", "email", "created_at", "birthdate"}).
// 			AddRow(expectedUser.ID, expectedUser.FullName, expectedUser.Username, expectedUser.Email, expectedUser.CreatedAt, expectedUser.Birthdate))

// 	// Call the GetUserByID function
// 	user, err := repo.GetUserByID(context.Background(), expectedUserID)
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 		return
// 	}

// 	// Verify the result
// 	if user == nil {
// 		t.Error("expected non-nil user, got nil")
// 		return
// 	}
// 	if user.ID != expectedUser.ID {
// 		t.Errorf("expected user ID %s, got %s", expectedUser.ID, user.ID)
// 	}
// 	if user.FullName != expectedUser.FullName {
// 		t.Errorf("expected user full name %s, got %s", expectedUser.FullName, user.FullName)
// 	}
// 	if user.Username != expectedUser.Username {
// 		t.Errorf("expected user username %s, got %s", expectedUser.Username, user.Username)
// 	}
// 	if user.Email != expectedUser.Email {
// 		t.Errorf("expected user email %s, got %s", expectedUser.Email, user.Email)
// 	}
// 	if !user.CreatedAt.Equal(expectedUser.CreatedAt) {
// 		t.Errorf("expected user created at %v, got %v", expectedUser.CreatedAt, user.CreatedAt)
// 	}
// 	if !user.Birthdate.Equal(expectedUser.Birthdate) {
// 		t.Errorf("expected user birthdate %v, got %v", expectedUser.Birthdate, user.Birthdate)
// 	}

// 	// Verify that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
