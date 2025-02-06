package subscriber

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type SubscriberViewer interface {
	ListNew(original *model.DiscordMessage, subs []*model.Subscriber) error
	Help(original *model.DiscordMessage) error
}
