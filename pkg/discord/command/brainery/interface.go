package brainery

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Commander interface {
	base.TextCommander

	Post(message *model.DiscordMessage) error
}
