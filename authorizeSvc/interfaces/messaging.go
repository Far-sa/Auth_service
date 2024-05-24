package interfaces

type EventPublisher interface {
	Publish(event string, data []byte) error
}
