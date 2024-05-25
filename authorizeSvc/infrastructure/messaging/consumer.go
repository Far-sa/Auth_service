package messaging

import (
	"authorization-service/infrastructure/messaging/rabbitmq"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConsumer handles consuming messages from RabbitMQ.

type RabbitMQConsumer struct {
	rabbitmqAdapter *rabbitmq.RabbitMQAdapter
	queueName       string
	routingKey      string
	exchange        string
	messageHandler  func(message string) error // Internal message handler
}

// NewRabbitMQConsumer creates a new RabbitMQConsumer with an internal message handler and sets up the queue and binding.
func NewRabbitMQConsumer(rabbitmqAdapter *rabbitmq.RabbitMQAdapter, queueName, routingKey, exchange string, handler func(message string) error) (*RabbitMQConsumer, error) {
	consumer := &RabbitMQConsumer{
		rabbitmqAdapter: rabbitmqAdapter,
		queueName:       queueName,
		routingKey:      routingKey,
		exchange:        exchange,
		messageHandler:  handler,
	}

	if err := consumer.setupQueueAndBinding(); err != nil {
		return nil, err
	}

	return consumer, nil
}

// setupQueueAndBinding sets up the queue and binding for the RabbitMQ consumer.
func (r *RabbitMQConsumer) setupQueueAndBinding() error {
	if _, err := r.rabbitmqAdapter.CreateQueue(r.queueName); err != nil {
		return fmt.Errorf("failed to create queue: %w", err)
	}

	if err := r.rabbitmqAdapter.CreateBinding(r.queueName, r.routingKey, r.exchange); err != nil {
		return fmt.Errorf("failed to create binding: %w", err)
	}

	return nil
}

// StartConsuming starts consuming messages from the configured queue.
func (r *RabbitMQConsumer) StartConsuming() error {
	msgs, err := r.rabbitmqAdapter.Consume(r.queueName)
	if err != nil {
		return fmt.Errorf("failed to start consuming messages: %w", err)
	}

	go func() {
		for d := range msgs {
			d := d // Create a local copy of the loop variable to avoid data races
			if err := r.handleMessage(d); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	return nil
}

func (r *RabbitMQConsumer) handleMessage(d amqp.Delivery) error {
	log.Printf("Received a message: %s", d.Body)
	if err := r.internalMessageHandler(string(d.Body)); err != nil {
		// Optionally, you could nack the message here if you want it to be requeued or logged elsewhere
		return fmt.Errorf("error processing message: %w", err)
	}

	// Acknowledge the message after successful processing
	if err := d.Ack(false); err != nil {
		return fmt.Errorf("error acknowledging message: %w", err)
	}

	return nil
}

// Example message handler
func (r *RabbitMQConsumer) internalMessageHandler(message string) error {
	log.Printf("Processing message: %s", message)
	// Implement your message processing logic here
	// For example, parse the message, process it, and store the results in a database

	return nil
}
