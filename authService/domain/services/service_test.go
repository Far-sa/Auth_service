package services

import (
	"authentication-service/domain/param"
	mocks "authentication-service/interfaces/mock"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	proto "shared/user"
)

type MockUserClient struct {
	mock.Mock
}

func (m *MockUserClient) GetUserByEmail(ctx context.Context, req *proto.GetUserEmailRequest) (*proto.GetUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*proto.GetUserResponse), args.Error(1)
}

func TestAuthService_Login(t *testing.T) {

	ctx := context.Background()

	type testCase struct {
		name          string
		loginRequest  param.LoginRequest
		setupMocks    func()
		expectedResp  param.LoginResponse
		expectedError error
	}

	cases := []testCase{
		{
			name: "successful login",
			loginRequest: param.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			expectedResp: param.LoginResponse{
				UserID: "1",
				Tokens: param.Token{
					AccessToken:  "access-token",
					RefreshToken: "refresh-token",
				},
			},
			expectedError: nil,
		},
		{
			name: "invalid credentials",
			loginRequest: param.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			expectedResp:  param.LoginResponse{},
			expectedError: errors.New("invalid credentials"),
		},
		{
			name: "user not found",
			loginRequest: param.LoginRequest{
				Email:    "notfound@example.com",
				Password: "password123",
			},
			expectedResp:  param.LoginResponse{},
			expectedError: errors.New("internal server error"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			mockAuthRepo := mocks.NewMockAuthRepository()
			mockUserClient := new(MockUserClient)
			authService := NewAuthService(mockAuthRepo, nil)

			resp, err := authService.Login(ctx, c.loginRequest)

			assert.Equal(t, c.expectedResp, resp)
			if c.expectedError != nil {
				assert.EqualError(t, err, c.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			mockAuthRepo.AssertExpectations(t)
			mockUserClient.AssertExpectations(t)
		})
	}
}
