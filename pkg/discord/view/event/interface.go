package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EventViewer interface {
	List(original *model.DiscordMessage, subs []*model.Event) error
}
