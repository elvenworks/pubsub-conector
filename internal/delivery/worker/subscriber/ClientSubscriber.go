package subscriber

import (
	"context"
	"sync"
	"time"

	pb "cloud.google.com/go/pubsub"
	"github.com/elvenworks/pubsub-conector/internal/driver/pubsub"
)

type ClientSubscriber struct {
	context *context.Context
	client  *pb.Client
}

func NewClientSubscriber(configs *pubsub.Config) (*ClientSubscriber, error) {
	client, err := pb.NewClient(configs.Context, configs.Credentials.ProjectID, configs.Option)

	if err != nil {
		return nil, err
	}

	return &ClientSubscriber{context: &configs.Context, client: client}, nil
}

func (c *ClientSubscriber) Subscription(topic string) (msg []byte, erro error) {
	sub := c.client.Subscription(topic)

	var mu sync.Mutex
	received := 0

	cctx, cancel := context.WithCancel(*c.context)

	err := sub.Receive(cctx, func(ctx context.Context, message *pb.Message) {
		mu.Lock()
		defer mu.Unlock()

		msg = message.Data

		message.Ack()
		received++
		if received == 1 {
			cancel()
		}
	})

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (c *ClientSubscriber) SubscriptionNack(topic string, timeout time.Duration) (success bool, erro error) {
	sub := c.client.Subscription(topic)

	var mu sync.Mutex
	received := 0

	cctx, cancel := context.WithTimeout(*c.context, time.Second*timeout)

	cm := make(chan *pb.Message)
	defer close(cm)

	err := sub.Receive(cctx, func(ctx context.Context, message *pb.Message) {
		mu.Lock()
		defer mu.Unlock()

		message.Nack()

		received++
		success = true
		if received == 1 {
			cancel()
		}

	})

	if err != nil {
		return false, err
	}

	return success, nil
}
