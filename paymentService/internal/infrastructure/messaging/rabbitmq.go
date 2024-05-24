package messaging

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn: conn, channel: channel}, nil
}

func (r *RabbitMQ) Publish(message []byte, queueName string) error {
	_, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
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

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}

//!!
func NewRabbitMQConsumer(conn *amqp.Connection) (*RabbitMQConsumer, error) {
    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    q, err := ch.QueueDeclare(
        "order_created_queue", // name
        true,                  // durable
        false,                 // delete when unused
        false,                 // exclusive
        false,                 // no-wait
        nil,                   // arguments
    )
    if err != nil {
        return nil, err
    }

    err = ch.QueueBind(
        q.Name,           // queue name
        "order.created",  // routing key
        "order_exchange", // exchange
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
            // Process the message
        }
    }()

    return &RabbitMQConsumer{channel: ch}, nil
}