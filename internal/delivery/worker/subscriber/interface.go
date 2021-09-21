package subscriber

import "time"

type IClientSubscriber interface {
	Subscription(topic string) (msg []byte, erro error)
	SubscriptionNack(topic string, timeout time.Duration) (erro error)
}
