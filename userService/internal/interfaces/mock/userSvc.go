package mocks

import (
	"context"
	"user-service/internal/param"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetUser(ctx context.Context, userID string) (param.UserProfileResponse, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(param.UserProfileResponse), args.Error(1)
}

func (m *UserServiceMock) GetUserByEmail(ctx context.Context, email string) (param.UserInfo, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(param.UserInfo), args.Error(1)
}

func (m *UserServiceMock) Register(ctx context.Context, req param.RegisterRequest) (param.RegisterResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(param.RegisterResponse), args.Error(1)
}
