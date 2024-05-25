package messaging

import (
	"authorization-service/infrastructure/messaging/rabbitmq"
	"log"
)

type RabbitMQConsumer struct {
	rabbitmqAdapter *rabbitmq.RabbitMQAdapter
	handler         func(message string) error
}

func NewRabbitMQConsumer(rabbitmqAdapter *rabbitmq.RabbitMQAdapter, handler func(message string) error) (*RabbitMQConsumer, error) {
	consumer := &RabbitMQConsumer{
		rabbitmqAdapter: rabbitmqAdapter,
		handler:         handler,
	}

	err := consumer.setupQueueAndBinding("user_authenticated_queue", "user.authenticated", "auth_exchange")
	if err != nil {
		return nil, err
	}

	go consumer.startConsuming("user_authenticated_queue")
	return consumer, nil
}

func (r *RabbitMQConsumer) setupQueueAndBinding(queueName, routingKey, exchange string) error {
	_, err := r.rabbitmqAdapter.CreateQueue(queueName)
	if err != nil {
		return err
	}

	return r.rabbitmqAdapter.CreateBinding(queueName, routingKey, exchange)
}

func (r *RabbitMQConsumer) startConsuming(queueName string) {
	msgs, err := r.rabbitmqAdapter.Consume(queueName)
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		if err := r.handler(string(d.Body)); err != nil {
			log.Printf("Error handling message: %s", err)
		}
	}
}
