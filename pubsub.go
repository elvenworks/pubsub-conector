package pubsub

import (
	"errors"
	"time"

	"github.com/elvenworks/pubsub-conector/internal/delivery/worker/publisher"
	"github.com/elvenworks/pubsub-conector/internal/delivery/worker/subscriber"
	"github.com/elvenworks/pubsub-conector/internal/driver/pubsub"
	"github.com/sirupsen/logrus"
)

type Secret struct {
	JsonCredentials []byte
}

type Pubsub struct {
	config    *pubsub.Config
	publisher publisher.IPublisher
}

func InitPubsub(secret Secret) (p *Pubsub, err error) {
	config, err := pubsub.NewConfig(secret.JsonCredentials)

	if err != nil {
		logrus.Errorf("Failed to send message to pubsub, err: %s\n", err)
		return nil, err
	}

	return &Pubsub{
		config: config,
	}, nil
}

func (p *Pubsub) Publish(topic string, message []byte, attributes map[string]string) error {

	if p.publisher == nil {
		publisher, err := publisher.NewPublisher(p.config)

		if err != nil {
			return err
		}

		p.publisher = publisher
	}

	err := p.publisher.Publish(topic, message, attributes)

	if err != nil {
		return err
	}

	return nil

}

func (p *Pubsub) PublishAndSubscriptionOnce(topic string, message []byte) error {
	publisher, err := publisher.NewPublisher(p.config)
	if err != nil {
		return err
	}

	err = publisher.Publish(topic, message, nil)
	if err != nil {
		return err
	}

	clientSubscriber, err := subscriber.NewClientSubscriber(p.config)
	if err != nil {
		return err
	}

	msg, err := clientSubscriber.Subscription(topic)
	if err != nil {
		return err
	}

	if msg == nil {
		return errors.New("any message has been consumed")
	}

	logrus.Info("Message consumed right after being produced: ", msg)

	return nil
}

func (p *Pubsub) SubscriptionNack(topic string, timeout time.Duration) error {
	clientSubscriber, err := subscriber.NewClientSubscriber(p.config)
	if err != nil {
		return err
	}

	success, err := clientSubscriber.SubscriptionNack(topic, timeout)
	if err != nil {
		return err
	}

	if !success {
		return errors.New("any message has been consumed")
	}

	return nil

}
