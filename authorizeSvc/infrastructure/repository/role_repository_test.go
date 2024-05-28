package repository_test

import (
	"authorization-service/infrastructure/repository"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestAssignRole tests the AssignRole method of DB.
func TestAssignRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.DB{}
	ctx := context.TODO()
	username := "testuser"
	role := "admin"

	// Prepare the expected SQL query
	query := `INSERT INTO roles \(username, role\) VALUES \(\$1, \$2\) ON CONFLICT \(username\) DO UPDATE SET role = \$2`
	mock.ExpectExec(query).WithArgs(username, role).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the method under test
	err = r.AssignRole(ctx, username, role)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
// TestCheckPermission tests the CheckPermission method of DB.
func TestCheckPermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.DB{}
	ctx := context.TODO()
	username := "testuser"
	permission := "read"
	// Prepare the expected SQL query
	query := `SELECT EXISTS\(SELECT 1 FROM permissions WHERE username = \$1 AND permission = \$2\)`
	mock.ExpectQuery(query).WithArgs(username, permission).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// Call the method under test
	hasPermission, err := r.CheckPermission(ctx, username, permission)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if !hasPermission {
		t.Errorf("expected permission to be true, but got false")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}// TestUpdateUserRoles tests the UpdateUserRoles method of DB.
func TestUpdateUserRoles(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.DB{}
	ctx := context.TODO()
	userID := "testuser"
	role := "admin"

	// Call the method under test
	err = r.UpdateUserRoles(ctx, userID, role)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Add assertions here if needed
}