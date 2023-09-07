package brainery

import (
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Brainery struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) Service {
	return &Brainery{
		adapter: adapter,
		l:       l,
	}
}

type PostInput struct {
	URL         string
	Author      string
	Reward      string
	Description string
	PublishedAt *time.Time
	Tags        []string
	Github      string
	DiscordID   string
}

type Service interface {
	Post(in *PostInput) (*model.Brainery, error)
	Report(view string, date string) (*model.BraineryMetric, error)
}
