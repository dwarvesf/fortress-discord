package event

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EventServicer interface {
	GetUpcomingEvents() ([]*model.Event, error)
}
