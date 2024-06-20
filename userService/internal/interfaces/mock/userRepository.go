package mocks

import (
	"context"
	"user-service/internal/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) GetUserByID(ctx context.Context, userID string) (*entity.UserProfile, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*entity.UserProfile), args.Error(1)
}

func (m *UserRepositoryMock) FindUserByEmail(ctx context.Context, Email string) (*entity.UserProfile, error) {
	args := m.Called(ctx, Email)
	return args.Get(0).(*entity.UserProfile), args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(ctx context.Context, user *entity.UserProfile) (*entity.UserProfile, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*entity.UserProfile), args.Error(1)
}
