package icy

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Icy struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) EarnCommander {
	return &Icy{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Icy) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Icy().GetWeeklyDistribution()
	if err != nil {
		e.L.Error(err, "can't get list of weekly icy distribution")
		return err
	}

	// 2. render
	return e.view.Icy().List(message, data)
}
