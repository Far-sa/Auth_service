package interfaces

type Messaging interface {
	Publish(message []byte, queueName string) error
	Subscribe(queueName string, handler func(message []byte)) error
}
