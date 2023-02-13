package index

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type IndexCommander interface {
	base.TextCommander

	Search(message *model.DiscordMessage) error
}
