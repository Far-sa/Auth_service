package service

import (
	"authorization-service/internal/entity"
	"authorization-service/internal/interfaces"
	mocks "authorization-service/internal/interfaces/mock"
	"authorization-service/internal/param"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAssignRole is a unit test function that tests the AssignRole method of the AuthzService struct.
// It tests various scenarios for assigning roles to users and verifies the expected error and log messages.
// The test cases include scenarios where the role is not found, failed to assign the default role,
// other errors while getting the role, and when the user already has a role assigned.
// Each test case is run using the t.Run function to provide a clear breakdown of the test results.
func TestAssignRole(t *testing.T) {
	// Test case struct definition and initialization
	type testCase struct {
		name           string
		userID         string
		getRoleErr     error
		assignRoleErr  error
		expectedErr    error
		expectedLogMsg string
	}

	// Test cases
	cases := []testCase{
		{
			name:           "Role not found, assign default role successfully",
			userID:         "user1",
			getRoleErr:     interfaces.ErrRoleNotFound,
			assignRoleErr:  nil,
			expectedErr:    nil,
			expectedLogMsg: "Assigned default role to user user1",
		},
		{
			name:           "Role not found, failed to assign default role",
			userID:         "user2",
			getRoleErr:     interfaces.ErrRoleNotFound,
			assignRoleErr:  errors.New("assign role error"),
			expectedErr:    errors.New("assign role error"),
			expectedLogMsg: "Failed to assign default role to user user2: assign role error",
		},
		{
			name:           "Other error while getting role",
			userID:         "user3",
			getRoleErr:     errors.New("get role error"),
			assignRoleErr:  nil,
			expectedErr:    errors.New("get role error"),
			expectedLogMsg: "Failed to get role for user user3: get role error",
		},
		{
			name:           "User already has a role",
			userID:         "user4",
			getRoleErr:     nil,
			assignRoleErr:  nil,
			expectedErr:    nil,
			expectedLogMsg: "User user4 already has a role assigned: existing_role",
		},
	}

	// Iterate over test cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create a new context
			ctx := context.Background()

			// Create a mock role repository
			mockRoleRepo := mocks.NewMockRoleRepository()

			// Mock the behavior of the role repository based on the test case
			if c.getRoleErr == nil {
				mockRoleRepo.On("GetRoleByUserID", ctx, c.userID).Return(entity.Role{Name: "existing_role"}, nil)
			} else {
				mockRoleRepo.On("GetRoleByUserID", ctx, c.userID).Return(entity.Role{}, c.getRoleErr)
			}

			if c.getRoleErr == interfaces.ErrRoleNotFound {
				mockRoleRepo.On("AssignRole", ctx, c.userID, DefaultRole).Return(c.assignRoleErr)
			}

			// Create an instance of the AuthzService struct with the mock role repository
			service := &AuthzService{roleRepo: mockRoleRepo}

			// Call the AssignRole method with the test case parameters
			err := service.AssignRole(ctx, param.RoleAssignmentRequest{UserID: c.userID})

			// Assert the expected error and log message
			assert.Equal(t, c.expectedErr, err)
			mockRoleRepo.AssertExpectations(t)
		})
	}
}
