package mocks

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

type MockUserEvents struct {
	mock.Mock
}

func NewUserEventMock() *MockUserEvents {
	return &MockUserEvents{}
}

func (m *MockUserEvents) DeclareExchange(name, kind string) error {
	args := m.Called(name, kind)
	return args.Error(0)
}

func (m *MockUserEvents) Publish(ctx context.Context, exchange, key string, msg amqp.Publishing) error {
	args := m.Called(ctx, exchange, key, msg)
	return args.Error(0)
}

func (m *MockUserEvents) CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error) {
	args := m.Called(queueName, durable, autodelete)
	return args.Get(0).(amqp.Queue), args.Error(1)
}

func (m *MockUserEvents) CreateBinding(queueName, routingKey, exchangeName string) error {
	args := m.Called(queueName, routingKey, exchangeName)
	return args.Error(0)
}

func (m *MockUserEvents) Consume(queueName, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	args := m.Called(queueName, consumer, autoAck)
	return args.Get(0).(<-chan amqp.Delivery), args.Error(1)
}
