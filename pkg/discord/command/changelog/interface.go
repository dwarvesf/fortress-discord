package changelog

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ChangelogCommander interface {
	base.TextCommander

	Send(message *model.DiscordMessage) error
}
