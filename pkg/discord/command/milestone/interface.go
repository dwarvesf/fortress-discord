package milestone

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MilestoneCommander interface {
	base.TextCommander

	ListMilestones(message *model.DiscordMessage) error
}
