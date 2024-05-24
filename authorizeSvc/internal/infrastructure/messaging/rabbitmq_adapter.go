package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQConnection(url string) (*RabbitMQConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to RabbitMQ successfully")

	return &RabbitMQConnection{conn, channel}, nil
}

func (r *RabbitMQConnection) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

// TODO consume method

// type RabbitMQPublisher struct {
// 	connection *RabbitMQConnection
// }

// func NewRabbitMQPublisher(connection *RabbitMQConnection) *RabbitMQPublisher {
// 	return &RabbitMQPublisher{connection}
// }

// func (p *RabbitMQPublisher) Publish(event string, data []byte) error {
// 	return p.connection.channel.Publish(
// 		"",    // exchange
// 		event, // routing key
// 		false, // mandatory
// 		false, // immediate
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        data,
// 		},
// 	)
// }

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
