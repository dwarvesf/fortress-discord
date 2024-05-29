package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EventCommander interface {
	base.TextCommander

	List(message *model.DiscordMessage) error
	ListGuildScheduledEvents(message *model.DiscordMessage) error
	SetSpeakers(message *model.DiscordMessage) error
}
