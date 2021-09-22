package pubsub

import "time"

type IPubsub interface {
	Publish(topic string, message []byte, attributes map[string]string) error
	PublishAndSubscriptionOnce(topic string, message []byte) error
	SubscriptionNack(topic string, timeout time.Duration) error
}
