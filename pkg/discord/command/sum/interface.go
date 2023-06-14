package sum

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type SumCommander interface {
	base.TextCommander

	Sum(message *model.DiscordMessage) error
}
