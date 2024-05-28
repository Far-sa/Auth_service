package migrator_test

import (
	"authorization-service/infrastructure/database/migrator"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockMigrate is a mock implementation of MigrateInterface.
type MockMigrate struct {
	mock.Mock
}

func (mm *MockMigrate) Up() error {
	args := mm.Called()
	return args.Error(0)
}

func (mm *MockMigrate) Down() error {
	args := mm.Called()
	return args.Error(0)
}

func TestNewMigrator(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://root:password@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer db.Close()

	migrationFilesPath := "/path/to/migrations"

	migrator, err := migrator.NewMigrator(db, migrationFilesPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if migrator == nil {
		t.Error("expected non-nil Migrator, got nil")
	}
}

// TestMigrator_Up tests the Up method of Migrator.
func TestMigrator_Up(t *testing.T) {
	mockMigrate := new(MockMigrate)
	migrator := migrator.Migrator{}

	// Set expectation for Up method
	mockMigrate.On("Up").Return(nil)

	// Call the method under test
	err := migrator.Up()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Assert that the expectations were met
	mockMigrate.AssertExpectations(t)
}

// TestMigrator_Down tests the Down method of Migrator.
func TestMigrator_Down(t *testing.T) {
	mockMigrate := new(MockMigrate)
	migrator := migrator.Migrator{}

	// Set expectation for Down method
	mockMigrate.On("Down").Return(nil)

	// Call the method under test
	err := migrator.Down()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Assert that the expectations were met
	mockMigrate.AssertExpectations(t)
}
