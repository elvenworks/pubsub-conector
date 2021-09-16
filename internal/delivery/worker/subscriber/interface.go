package subscriber

type IClientSubscriber interface {
	Subscription(topic string) (msg []byte, erro error)
	GetLag(topic string) (lagTotal int64, err error)
}
