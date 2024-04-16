package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MemoCommander interface {
	base.TextCommander

	List(message *model.DiscordMessage) error
	Sync(message *model.DiscordMessage) error
}
