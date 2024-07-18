package researchtopic

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/constant"
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

func (e *ResearchTopic) GetDiscordResearchTopics(timeRange string) (*model.DiscordResearchTopicResponse, error) {
	if timeRange == constant.AllTime {
		timeRange = "0"
	}

	// get response from fortress
	adapterResearchTopic, err := e.adapter.Fortress().GetDiscordResearchTopics(timeRange)
	if err != nil {
		e.l.Error(err, "can't get memo logs from fortress")
		return nil, err
	}

	return adapterResearchTopic, nil
}
