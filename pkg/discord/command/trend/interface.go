package trend

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type TrendCommander interface {
	base.TextCommander
	DefaultTrend(message *model.DiscordMessage) error
	Trend(message *model.DiscordMessage) error
}
