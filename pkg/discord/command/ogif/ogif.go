package ogif

import (
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c command) FetchOgifStats(msg *model.DiscordMessage, discordID string, after time.Time, timeAmount int, timeUnit string) error {
	logger := c.L.AddField("discordID", discordID).AddField("after", after)
	stats, err := c.svc.Event().GetOgifStats(discordID, after)
	if err != nil {
		logger.Error(err, "error when get ogif stats")
		return err
	}

	return c.view.Ogif().RenderOgifStats(msg, discordID, stats, timeAmount, timeUnit)
}

func (c command) GetOgifLeaderboard(msg *model.DiscordMessage, after time.Time, timeAmount int, timeUnit string) error {
	logger := c.L.AddField("after", after)
	leaderboard, err := c.svc.Event().GetOgifLeaderboard(after, 10)
	if err != nil {
		logger.Error(err, "error when get ogif leaderboard")
		return err
	}

	return c.view.Ogif().RenderOgifLeaderboard(msg, leaderboard, timeAmount, timeUnit)
}
