package sum

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Sum struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) SumCommander {
	return &Sum{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Sum) Sum(message *model.DiscordMessage) error {
	url := message.ContentArgs[1]
	// 1. get data from service
	data, err := e.svc.Sum().SummarizeArticle(url)
	if err != nil {
		e.L.Error(err, "failed to summarize the given article")
		return err
	}

	// 2. render
	return e.view.Sum().Sum(message, data)
}
