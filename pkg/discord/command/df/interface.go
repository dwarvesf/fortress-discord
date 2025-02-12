package df

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DFCommander interface {
	base.TextCommander

	ProcessWithN8N(message *model.DiscordMessage) error
}
