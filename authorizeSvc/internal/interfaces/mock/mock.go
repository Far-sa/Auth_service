package mocks

import (
	"authorization-service/internal/entity"
	"authorization-service/internal/param"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

type MockAuthorizationService struct {
	mock.Mock
}

func NewMockAuthorizationService() *MockAuthorizationService {
	return &MockAuthorizationService{}
}

func (m *MockAuthorizationService) AssignRole(ctx context.Context, userID param.RoleAssignmentRequest) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockAuthorizationService) UpdateUserRole(ctx context.Context, newRole param.RoleUpdateRequest) error {
	args := m.Called(ctx, newRole)
	return args.Error(0)
}

type MockRoleRepository struct {
	mock.Mock
}

func NewMockRoleRepository() *MockRoleRepository {
	return &MockRoleRepository{}
}

func (m *MockRoleRepository) CheckPermission(ctx context.Context, username, permission string) (bool, error) {
	args := m.Called(ctx, username, permission)
	return args.Bool(0), args.Error(1)
}

func (m *MockRoleRepository) AssignRole(ctx context.Context, userID, role string) error {
	args := m.Called(ctx, userID, role)
	return args.Error(0)
}

func (m *MockRoleRepository) GetRoleByUserID(ctx context.Context, userID string) (entity.Role, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(entity.Role), args.Error(1)
}

func (m *MockRoleRepository) UpdateRole(ctx context.Context, userID string, newRole entity.Role) error {
	args := m.Called(ctx, userID, newRole)
	return args.Error(0)
}

type MockRoleEvents struct {
	mock.Mock
}

func NewMockRoleEvents() *MockRoleEvents {
	return &MockRoleEvents{}
}

func (m *MockRoleEvents) DeclareExchange(name, kind string) error {
	args := m.Called(name, kind)
	return args.Error(0)
}

func (m *MockRoleEvents) CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error) {
	args := m.Called(queueName, durable, autodelete)
	return args.Get(0).(amqp.Queue), args.Error(1)
}

func (m *MockRoleEvents) CreateBinding(queueName, routingKey, exchangeName string) error {
	args := m.Called(queueName, routingKey, exchangeName)
	return args.Error(0)
}

func (m *MockRoleEvents) Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	args := m.Called(queueName, consumer, autoAck)
	return args.Get(0).(<-chan amqp.Delivery), args.Error(1)
}

func (m *MockRoleEvents) Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	args := m.Called(ctx, exchange, routingKey, options)
	return args.Error(0)
}
