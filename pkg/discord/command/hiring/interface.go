package hiring

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type HiringCommander interface {
	base.TextCommander

	OpenPositions(message *model.DiscordMessage) error
}
