package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// MessagePublisher defines methods for publishing messages

type RabbitMQPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQPublisher(connStr string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(
		"auth_exchange", // Name of the exchange
		"topic",         // Exchange type (fanout for broadcasting)
		false,           // Durable (survives server restarts)
		false,           // Delete when unused
		false,           // Internal
		false,           // No wait
		nil,             // Arguments
	)
	if err != nil {
		return nil, err
	}
	return &RabbitMQPublisher{conn: conn, ch: ch}, nil
}

func (p *RabbitMQPublisher) PublishUserAuthenticated(userID string) error {
	body := fmt.Sprintf("UserAuthenticated: %s", userID)
	err := p.ch.Publish(
		"auth_exchange",      // exchange
		"user.authenticated", // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s", body)
	return nil
}

func (p *RabbitMQPublisher) Publish(ctx context.Context, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = p.ch.PublishWithContext(ctx, "", "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	return err
}

func (p *RabbitMQPublisher) Close() error {
	if p.ch != nil {
		if err := p.ch.Close(); err != nil {
			return err
		}
	}
	if p.conn != nil {
		return p.conn.Close()
	}
	return nil
}
