package publisher

import (
	"context"

	pb "cloud.google.com/go/pubsub"
	"github.com/elvenworks/pubsub-conector/internal/driver/pubsub"
)

type Publisher struct {
	context   *context.Context
	publisher *pb.Client
}

func NewPublisher(configs *pubsub.Config) (*Publisher, error) {
	publisher, err := pb.NewClient(configs.Context, configs.Credentials.ProjectID, configs.Option)

	if err != nil {
		return nil, err
	}

	return &Publisher{
		context:   &configs.Context,
		publisher: publisher,
	}, nil
}

func (p *Publisher) Publish(topic string, message []byte, attributes map[string]string) (err error) {
	id := p.publisher.Topic(topic)

	result := id.Publish(*p.context, &pb.Message{
		Data:       []byte(message),
		Attributes: attributes,
	})

	_, err = result.Get(*p.context)

	if err != nil {
		return err
	}

	err = p.Close()
	if err != nil {
		return err
	}

	return nil
}

func (p *Publisher) Close() error {
	if p != nil {
		return p.publisher.Close()
	}
	return nil
}
