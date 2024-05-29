package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EventViewer interface {
	Help(original *model.DiscordMessage) error
	List(original *model.DiscordMessage, subs []*model.Event) error
	ListScheduledEvents(original *model.DiscordMessage, subs []*model.Event) error
}
