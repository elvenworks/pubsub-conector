package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/elvenworks/pubsub-conector/internal/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type Config struct {
	Context     context.Context
	Credentials domain.Credentials
	Option      option.ClientOption
}

func NewConfig(jsonCredentials []byte) (c *Config, err error) {
	var context = context.Background()
	var credentials domain.Credentials

	if err := json.Unmarshal(jsonCredentials, &credentials); err != nil {
		logrus.Errorf("Failed to unmarshal credentials to pubsub, err: %s\n", err)
		return nil, err
	}

	if credentials.ProjectID == "" {
		logrus.Errorf("ProjectID not found to pubsub.\n")
		return nil, err
	}

	creds, err := google.CredentialsFromJSON(context, jsonCredentials, pubsub.ScopePubSub)

	if err != nil {
		return nil, err
	}

	config := &Config{Context: context, Credentials: credentials, Option: option.WithCredentials(creds)}

	return config, nil

}
