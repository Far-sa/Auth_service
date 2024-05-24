package messaging

import (
	"authentication-service/interfaces"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	channel *amqp.Channel
}

func NewRabbitMQConnection(url string) (*RabbitMQConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQConnection{channel: ch}, nil
}

func (c *RabbitMQConnection) GetChannel() *amqp.Channel {
	return c.channel
}

func (c *RabbitMQConnection) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
}

type RabbitMQPublisher struct {
	connection *RabbitMQConnection
}

func NewRabbitMQPublisher(connection *RabbitMQConnection) interfaces.EventPublisher {
	return &RabbitMQPublisher{connection: connection}
}

func (p *RabbitMQPublisher) Publish(event string, message []byte) error {
	return p.connection.GetChannel().Publish(
		"",    // exchange
		event, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
}

func SetupRabbitMQ(conn *amqp.Connection) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// Declare an exchange
	err = channel.ExchangeDeclare(
		"exchange_name", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return err
	}

	// Declare a queue
	_, err = channel.QueueDeclare(
		"queue_name", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	// Bind the queue to the exchange
	err = channel.QueueBind(
		"queue_name",    // queue name
		"routing_key",   // routing key
		"exchange_name", // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("RabbitMQ setup completed successfully.")
	return nil
}
