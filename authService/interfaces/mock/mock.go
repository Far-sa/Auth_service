package mocks

import (
	"authentication-service/domain/entities"
	"authentication-service/domain/param"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

type MockAuthenticationService struct {
	mock.Mock
}

func NewMockAuthenticationService() *MockAuthenticationService {
	return &MockAuthenticationService{}
}

func (m *MockAuthenticationService) Login(ctx context.Context, loginRequest param.LoginRequest) (param.LoginResponse, error) {
	args := m.Called(ctx, loginRequest)
	return args.Get(0).(param.LoginResponse), args.Error(1)
}

// ... implement other mock authentication methods

type MockAuthRepository struct {
	mock.Mock
}

func NewMockAuthRepository() *MockAuthRepository {
	return &MockAuthRepository{}
}

func (m *MockAuthRepository) SaveToken(ctx context.Context, tokens *entities.Token) error {
	args := m.Called(ctx, tokens)
	return args.Error(0)
}

type MockAuthEvents struct {
	mock.Mock
}

func NewMockAuthEvents() *MockAuthEvents {
	return &MockAuthEvents{}
}

func (m *MockAuthEvents) DeclareExchange(name, kind string) error {
	args := m.Called(name, kind)
	return args.Error(0)
}

func (m *MockAuthEvents) CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error) {
	args := m.Called(queueName, durable, autodelete)
	return args.Get(0).(amqp.Queue), args.Error(1)
}

func (m *MockAuthEvents) CreateBinding(queueName, routingKey, exchangeName string) error {
	args := m.Called(queueName, routingKey, exchangeName)
	return args.Error(0)
}

func (m *MockAuthEvents) Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	args := m.Called(queueName, consumer, autoAck)
	return args.Get(0).(<-chan amqp.Delivery), args.Error(1)
}

func (m *MockAuthEvents) Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	args := m.Called(ctx, exchange, routingKey, options)
	return args.Error(0)
}
