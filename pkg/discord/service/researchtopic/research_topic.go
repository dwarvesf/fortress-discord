package researchtopic

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ResearchTopic struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) ResearchTopicServicer {
	return &ResearchTopic{
		adapter: adapter,
		l:       l,
	}
}

func (e *ResearchTopic) GetDiscordResearchTopics(page, size string) (*model.DiscordResearchTopicResponse, error) {
	// get response from fortress
	adapterResearchTopic, err := e.adapter.Fortress().GetDiscordResearchTopics(page, size)
	if err != nil {
		e.l.Error(err, "can't get memo logs from fortress")
		return nil, err
	}

	return adapterResearchTopic, nil
}
