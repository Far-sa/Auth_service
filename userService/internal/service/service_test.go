package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"user-service/internal/entity"
	mocks "user-service/internal/interfaces/mock"
	"user-service/internal/param"
	"user-service/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	mockRepo := mocks.NewUserRepositoryMock()
	mockEvents := mocks.NewUserEventMock()

	ctx := context.Background()

	req := param.RegisterRequest{
		UserName:    "testuser",
		Email:       "test@example.com",
		Password:    "password123",
		FullName:    "Test User",
		PhoneNumber: "1234567890",
	}

	hashedPassword, _ := utils.HashPassword(req.Password)
	expectedUser := &entity.UserProfile{
		ID:          "1",
		Username:    req.UserName,
		Email:       req.Email,
		Password:    hashedPassword,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now(),
	}

	mockRepo.On("CreateUser", ctx, mock.AnythingOfType("*entity.UserProfile")).Return(expectedUser, nil).Once()
	mockEvents.On("DeclareExchange", "user_events_exchange", "topic").Return(nil)
	mockEvents.On("Publish", ctx, "user_events_exchange", "user.created", mock.AnythingOfType("amqp.Publishing")).Return(nil)

	userSvc := NewUserService(mockRepo, nil)
	resp, err := userSvc.Register(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, resp.User.ID)
	assert.Equal(t, expectedUser.Email, resp.User.Email)
	assert.Equal(t, expectedUser.FullName, resp.User.FullName)
	assert.Equal(t, expectedUser.PhoneNumber, resp.User.PhoneNumber)

	mockRepo.AssertExpectations(t)
	mockEvents.AssertExpectations(t)
}

func TestGetUserByEmail(t *testing.T) {
	mockRepo := mocks.NewUserRepositoryMock()
	userSvc := NewUserService(mockRepo, nil)

	ctx := context.Background()
	email := "test@example.com"

	t.Run("user found", func(t *testing.T) {
		expectedUser := &entity.UserProfile{
			ID:       "1",
			Email:    email,
			FullName: "Test User",
		}

		mockRepo.On("FindUserByEmail", ctx, email).Return(expectedUser, nil).Once()

		resp, err := userSvc.GetUserByEmail(ctx, email)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, resp.ID)
		assert.Equal(t, expectedUser.Email, resp.Email)
		assert.Equal(t, expectedUser.FullName, resp.FullName)

		mockRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("FindUserByEmail", ctx, email).Return(nil, errors.New("user not found"))

		resp, err := userSvc.GetUserByEmail(ctx, email)

		assert.Error(t, err)
		assert.Empty(t, resp.ID)
		assert.Empty(t, resp.Email)
		assert.Empty(t, resp.FullName)

		mockRepo.AssertExpectations(t)
	})
}
func TestGetUser(t *testing.T) {
	mockRepo := mocks.NewUserRepositoryMock()
	userSvc := NewUserService(mockRepo, nil)

	ctx := context.Background()
	userID := "1"
	expectedUser := &entity.UserProfile{
		ID:        userID,
		Email:     "test@example.com",
		FullName:  "Test User",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetUserByID", ctx, userID).Return(expectedUser, nil).Once()

	resp, err := userSvc.GetUser(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, resp.UserProfile.ID)
	assert.Equal(t, expectedUser.Email, resp.UserProfile.Email)
	assert.Equal(t, expectedUser.FullName, resp.UserProfile.Username)
	assert.Equal(t, expectedUser.CreatedAt, resp.UserProfile.CreatedAt)
	mockRepo.AssertExpectations(t)
}
