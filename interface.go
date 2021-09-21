package pubsub

import "time"

type IPubsub interface {
	Publish(topic string, message []byte, attributes map[string]string) error
	PublishAndSubscriptionOnce(topic string, message []byte) error
	SubscriptionNack(topic string, message []byte, timeout time.Duration) error
}
