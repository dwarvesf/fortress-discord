package event

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EventServicer interface {
	GetUpcomingEvents() ([]*model.Event, error)
	GetGuildScheduledEvents() ([]*model.DiscordEvent, error)
	CreateGuildScheduledEvent(*model.DiscordEvent) error
	SetSpeakers(message *model.DiscordMessage) error
}
