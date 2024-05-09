package event

import (
	"fmt"

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

func (e *Event) ListGuildScheduledEvents(message *model.DiscordMessage) error {
	// 1. get data from service
	events, err := e.svc.Event().GetGuildScheduledEvents()
	if err != nil {
		e.L.Error(err, "can't get list of discord scheduled events")
		return err
	}

	var data []*model.Event
	for _, ev := range events {
		data = append(data, &model.Event{
			Id:   ev.DiscordEventID,
			Name: ev.Name,
			Date: model.EventDate{
				Time:    &ev.Date,
				HasTime: true,
			},
		})
	}

	// 2. render
	return e.view.Event().ListScheduledEvents(message, data)
}

func (e *Event) SetSpeakers(message *model.DiscordMessage) error {
	err := e.svc.Event().SetSpeakers(message)
	if err != nil {
		fmt.Println("error", err)
		e.L.Error(err, "can't set speakers for the scheduled event")
		return e.view.Done().Error(message, err.Error())
	}
	return e.view.Done().Success(message)
}
