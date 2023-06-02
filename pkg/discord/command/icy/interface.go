package icy

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EarnCommander interface {
	base.TextCommander

	List(message *model.DiscordMessage) error
}
