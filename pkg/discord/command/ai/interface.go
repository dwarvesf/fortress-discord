package ai

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type AICommander interface {
	base.TextCommander

	ProcessAI(message *model.DiscordMessage) error
}
