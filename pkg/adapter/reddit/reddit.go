package reddit

import (
	"fmt"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type adapter struct {
	client *reddit.Client
}

func New(cfg *config.Config, l logger.Logger) (Adapter, error) {
	clientID := cfg.Reddit.ClientID
	if clientID == "" {
		l.Info("reddit client id is empty")
	}

	clientSecret := cfg.Reddit.ClientSecret
	if clientSecret == "" {
		l.Info("reddit client secret is empty")
	}

	username := cfg.Reddit.Username
	if username == "" {
		l.Info("reddit username is empty")
	}

	password := cfg.Reddit.Password
	if password == "" {
		l.Info("reddit password is empty")
	}

	auth := reddit.Credentials{
		ID:       clientID,
		Secret:   clientSecret,
		Username: username,
		Password: password,
	}

	client, err := reddit.NewClient(auth, reddit.WithUserAgent("fortress-bot"))
	if err != nil {
		return nil, fmt.Errorf("create reddit client failed: %w", err)
	}

	return &adapter{
		client: client,
	}, nil
}
