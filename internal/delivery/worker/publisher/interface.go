package publisher

type IPublisher interface {
	Publish(topic string, message []byte, attributes map[string]string) error
	Close() error
}
