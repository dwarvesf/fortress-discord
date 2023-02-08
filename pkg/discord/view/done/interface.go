package done

import "github.com/dwarvesf/fortress-discord/pkg/model"

type DoneViewer interface {
	Repost(original *model.DiscordMessage, msg string, channel string, icy string) error

	MissingContent(original *model.DiscordMessage) error
	CantSendReward(original *model.DiscordMessage) error
}
