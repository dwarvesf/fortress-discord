package topic

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Topic struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) TopicCommander {
	return &Topic{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

func (e *Topic) List(message *model.DiscordMessage) error {
	page := "1"
	size := e.cfg.DiscordResearchTopic.Size

	data, err := e.svc.ResearchTopic().GetDiscordResearchTopics(page, size)
	if err != nil {
		e.L.Error(err, "can't get list of discord research topic")
		return err
	}

	// 2. render
	return e.view.Topic().List(message, page, size, *data)
}
