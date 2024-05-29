package messaging

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Enable publisher confirms (optional)
	if err := ch.Confirm(false); err != nil {
		ch.Close()   // Close the channel on Confirm error
		conn.Close() // Also close the connection
		return nil, fmt.Errorf("failed to enable publisher confirms: %w", err)
	}

	return &RabbitMQ{conn: conn, ch: ch}, nil
}

// CreateExchange declares a new exchange on the RabbitMQ server
func (rc RabbitMQ) DeclareExchange(name, kind string) error {

	return rc.ch.ExchangeDeclare(
		name,  // Name of the exchange
		kind,  // Type of exchange (e.g., "fanout", "direct", "topic")
		true,  // Durable (survives server restarts)
		false, // Delete when unused
		false, // Exclusive (only this connection can access)
		false,
		nil, // Arguments
	)
}

func (rc RabbitMQ) CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error) {
	q, err := rc.ch.QueueDeclare(queueName, durable, autodelete, false, false, nil)
	if err != nil {
		return amqp.Queue{}, nil
	}
	return q, err
}

func (rc RabbitMQ) CreateBinding(name, binding, exchange string) error {
	return rc.ch.QueueBind(name, binding, exchange, false, nil)
}

// ! PublishMessage sends a message to a specific exchange with a routing key
func (rc RabbitMQ) Publish(ctx context.Context, exchangeName string, routingKey string, options amqp.Publishing) error {

	return rc.ch.PublishWithContext(
		ctx,
		exchangeName, // Name of the exchange
		routingKey,   // Routing key for message
		false,        // Mandatory (if true, message is rejected if no queue is bound)
		false,        // Immediate (if true, delivery happens now, or fails)
		options,
	)

}

func (rc RabbitMQ) Consume(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return rc.ch.Consume(queue, consumer, autoAck, false, false, false, nil)
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}
