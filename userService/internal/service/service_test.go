package service

import (
	"context"
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

	mockRepo.On("CreateUser", ctx, mock.AnythingOfType("*entity.UserProfile")).Return(expectedUser, nil)
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
