package messaging

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	exchange   string
}

func NewRabbitMQ(url, exchange string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchange, // name of the exchange
		"topic",  // type
		true,     // durable
		false,    // delete when unused
		false,    // internal
		false,    // noWait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		connection: conn,
		channel:    ch,
		exchange:   exchange,
	}, nil
}

func (r *RabbitMQ) Publish(routingKey string, body []byte) error {
	err := r.channel.Publish(
		r.exchange, // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	return err
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.connection.Close()
}
