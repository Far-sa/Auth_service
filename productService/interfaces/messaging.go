package interfaces

type Messaging interface {
	Publish(routingKey string, body []byte) error
	Close()
}
