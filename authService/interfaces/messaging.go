package interfaces

type EventPublisher interface {
	Publish(event string, message []byte) error
}
