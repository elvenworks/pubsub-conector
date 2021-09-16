package pubsub

type IPubsub interface {
	InitPubsub(secret Secret) *Pubsub
	Publish(topic string, message []byte, attributes map[string]string) error
	PublishAndSubscriptionOnce(topic string, message []byte) error
	GetLag(topic string) (lagTotal int64, err error)
}
