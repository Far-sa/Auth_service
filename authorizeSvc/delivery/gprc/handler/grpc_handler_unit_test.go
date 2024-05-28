package handler_test

import (
	"authorization-service/delivery/gprc/handler"
	"authorization-service/pb"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthorizationService is a mock implementation of interfaces.AuthorizationService
type MockAuthorizationService struct {
	mock.Mock
}

func (m *MockAuthorizationService) AssignRole(ctx context.Context, username string, role string) error {
	args := m.Called(ctx, username, role)
	return args.Error(0)
}

func (m *MockAuthorizationService) CheckPermission(ctx context.Context, username string, permission string) (bool, error) {
	args := m.Called(ctx, username, permission)
	return args.Bool(0), args.Error(1)
}

func TestAssignRole(t *testing.T) {
	mockAuthService := new(MockAuthorizationService)
	handler := handler.NewAuthzHandler(mockAuthService)

	testCases := []struct {
		name       string
		username   string
		role       string
		mockError  error
		expectResp *pb.AssignRoleResponse
		expectErr  error
	}{
		{
			name:       "Success",
			username:   "testuser",
			role:       "admin",
			mockError:  nil,
			expectResp: &pb.AssignRoleResponse{Message: "Role assigned successfully"},
			expectErr:  nil,
		},
		{
			name:       "Error",
			username:   "testuser",
			role:       "admin",
			mockError:  errors.New("assignment failed"),
			expectResp: nil,
			expectErr:  errors.New("assignment failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.AssignRoleRequest{Username: tc.username, Role: tc.role}
			mockAuthService.On("AssignRole", mock.Anything, tc.username, tc.role).Return(tc.mockError)

			resp, err := handler.AssignRole(context.Background(), req)

			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectResp, resp)
			}
			mockAuthService.AssertExpectations(t)
		})
	}
}

func TestCheckPermission(t *testing.T) {
	mockAuthService := new(MockAuthorizationService)
	handler := handler.NewAuthzHandler(mockAuthService)

	testCases := []struct {
		name       string
		username   string
		permission string
		mockResult bool
		mockError  error
		expectResp *pb.CheckPermissionResponse
		expectErr  error
	}{
		{
			name:       "PermissionGranted",
			username:   "testuser",
			permission: "write",
			mockResult: true,
			mockError:  nil,
			expectResp: &pb.CheckPermissionResponse{HasPermission: true},
			expectErr:  nil,
		},
		{
			name:       "PermissionDenied",
			username:   "testuser",
			permission: "write",
			mockResult: false,
			mockError:  nil,
			expectResp: &pb.CheckPermissionResponse{HasPermission: false},
			expectErr:  nil,
		},
		{
			name:       "Error",
			username:   "testuser",
			permission: "write",
			mockResult: false,
			mockError:  errors.New("permission check failed"),
			expectResp: nil,
			expectErr:  errors.New("permission check failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.CheckPermissionRequest{Username: tc.username, Permission: tc.permission}
			mockAuthService.On("CheckPermission", mock.Anything, tc.username, tc.permission).Return(tc.mockResult, tc.mockError)

			resp, err := handler.CheckPermission(context.Background(), req)

			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectResp, resp)
			}
			mockAuthService.AssertExpectations(t)
		})
	}
}

// func TestNewAuthzHandler(t *testing.T) {
// 	mockAuthService := new(MockAuthorizationService)
// 	handler := handler.NewAuthzHandler(mockAuthService)
// 	assert.NotNil(t, handler)
// 	assert.Equal(t, mockAuthService, handler.service)
// }

// func TestAssignRole(t *testing.T) {
// 	mockAuthService := new(MockAuthorizationService)
// 	handler := handler.NewAuthzHandler(mockAuthService)

// 	testUsername := "testuser"
// 	testRole := "admin"
// 	req := &pb.AssignRoleRequest{Username: testUsername, Role: testRole}

// 	mockAuthService.On("AssignRole", mock.Anything, testUsername, testRole).Return(nil)

// 	resp, err := handler.AssignRole(context.Background(), req)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	assert.Equal(t, "Role assigned successfully", resp.Message)
// }

// func TestAssignRole_Error(t *testing.T) {
// 	mockAuthService := new(MockAuthorizationService)
// 	handler := handler.NewAuthzHandler(mockAuthService)

// 	testUsername := "testuser"
// 	testRole := "admin"
// 	req := &pb.AssignRoleRequest{Username: testUsername, Role: testRole}
// 	testError := errors.New("assignment failed")

// 	mockAuthService.On("AssignRole", mock.Anything, testUsername, testRole).Return(testError)

// 	resp, err := handler.AssignRole(context.Background(), req)
// 	assert.Error(t, err)
// 	assert.Nil(t, resp)
// 	assert.Equal(t, testError, err)
// }

// func TestCheckPermission(t *testing.T) {
// 	mockAuthService := new(MockAuthorizationService)
// 	handler := handler.NewAuthzHandler(mockAuthService)

// 	testUsername := "testuser"
// 	testPermission := "write"
// 	req := &pb.CheckPermissionRequest{Username: testUsername, Permission: testPermission}

// 	mockAuthService.On("CheckPermission", mock.Anything, testUsername, testPermission).Return(true, nil)

// 	resp, err := handler.CheckPermission(context.Background(), req)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	assert.True(t, resp.HasPermission)
// }

// func TestCheckPermission_Error(t *testing.T) {
// 	mockAuthService := new(MockAuthorizationService)
// 	handler := handler.NewAuthzHandler(mockAuthService)

// 	testUsername := "testuser"
// 	testPermission := "write"
// 	req := &pb.CheckPermissionRequest{Username: testUsername, Permission: testPermission}
// 	testError := errors.New("permission check failed")

// 	mockAuthService.On("CheckPermission", mock.Anything, testUsername, testPermission).Return(false, testError)

// 	resp, err := handler.CheckPermission(context.Background(), req)
// 	assert.Error(t, err)
// 	assert.Nil(t, resp)
// 	assert.Equal(t, testError, err)
// }
