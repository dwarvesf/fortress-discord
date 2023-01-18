package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Event struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) EventServicer {
	return &Event{
		adapter: adapter,
		l:       l,
	}
}

func (e *Event) GetUpcomingEvents() ([]*model.Event, error) {
	// get response from fortress
	adapterEvents, err := e.adapter.Fortress().GetUpcomingEvents()
	if err != nil {
		e.l.Error(err, "can't get open positions from fortress")
		return nil, err
	}

	// normalized into in-app model
	events := adapterEvents.Data

	return events, nil
}
