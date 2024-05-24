package messaging

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
}

func NewRabbitMQ(rabbitMQURL, exchange string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:     conn,
		channel:  ch,
		exchange: exchange,
	}, nil
}

func (r *RabbitMQ) Publish(message []byte, queueName string) error {
	err := r.channel.Publish(
		r.exchange, // exchange
		queueName,  // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	return err
}

func (r *RabbitMQ) Subscribe(queueName string, handler func(message []byte)) error {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()
	return nil
}

// !
func NewRabbitMQPublisher(conn *amqp.Connection) (*RabbitMQPublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"order_exchange", // name of the exchange
		"topic",          // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{channel: ch}, nil
}

func (r *RabbitMQPublisher) PublishOrderCreated(orderID string) error {
	body := fmt.Sprintf("OrderCreated: %s", orderID)
	err := r.channel.Publish(
		"order_exchange", // exchange
		"order.created",  // routing key
		false,            // mandatory
		false,            // immediate
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
