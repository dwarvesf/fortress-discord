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
