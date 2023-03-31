package updates

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type UpdatesCommander interface {
	base.TextCommander

	List(message *model.DiscordMessage) error
	PreSend(message *model.DiscordMessage) error
}
