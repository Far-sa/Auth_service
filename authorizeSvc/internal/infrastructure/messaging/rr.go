package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	channel *amqp.Channel
}

func NewRabbitMQConsumer(conn *amqp.Connection) (*RabbitMQConsumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"user_authenticated_queue", // name
		true,                       // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,               // queue name
		"user.authenticated", // routing key
		"auth_exchange",      // exchange
		false,
		nil)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// Process the message and update user roles/permissions
		}
	}()

	return &RabbitMQConsumer{channel: ch}, nil
}
