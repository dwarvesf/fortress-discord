package salary

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type SalaryCommander interface {
	base.TextCommander

	Advance(message *model.DiscordMessage) error
}
