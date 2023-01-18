package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Event struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) EventCommander {
	return &Event{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Event) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Event().GetUpcomingEvents()
	if err != nil {
		e.L.Error(err, "can't get list of active earn")
		return err
	}

	// 2. render
	return e.view.Event().List(message, data)
}
