package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQAdapter(url string) (*RabbitMQAdapter, error) {
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

	return &RabbitMQAdapter{connection: conn, channel: ch}, nil
}

func (r *RabbitMQAdapter) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil {
		r.connection.Close()
	}
}

func (r *RabbitMQAdapter) DeclareExchange(name, kind string) error {

	return r.channel.ExchangeDeclare(
		name,  // Name of the exchange
		kind,  // Type of exchange (e.g., "fanout", "direct", "topic")
		true,  // Durable (survives server restarts)
		false, // Delete when unused
		false, // Exclusive (only this connection can access)
		false,
		nil, // Arguments
	)
}

func (r *RabbitMQAdapter) CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error) {
	return r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func (r *RabbitMQAdapter) CreateBinding(queueName, routingKey, exchange string) error {
	return r.channel.QueueBind(
		queueName,  // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
}

func (r *RabbitMQAdapter) Publish(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	return r.channel.PublishWithContext(
		ctx,
		exchange,   // Name of the exchange
		routingKey, // Routing key for message
		false,      // Mandatory (if true, message is rejected if no queue is bound)
		false,      // Immediate (if true, delivery happens now, or fails)
		options,
	)
}

func (r *RabbitMQAdapter) Consume(queueName string, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}
