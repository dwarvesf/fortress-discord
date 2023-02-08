package done

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DoneCommander interface {
	base.TextCommander

	Done(message *model.DiscordMessage) error
}
