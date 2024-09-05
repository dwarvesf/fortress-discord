package ogif

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type OgifViewer interface {
	RenderOgifStats(original *model.DiscordMessage, userID string, stats model.OgifStats, timeAmount int, timeUnit string) error
	RenderOgifLeaderboard(msg *model.DiscordMessage, leaderboard []model.OgifLeaderboardRecord, timeAmount int, timeUnit string) error
	Help(message *model.DiscordMessage) error
}
