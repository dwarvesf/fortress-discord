package digest

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DigestCommander interface {
	base.TextCommander

	List(message *model.DiscordMessage) error
}
