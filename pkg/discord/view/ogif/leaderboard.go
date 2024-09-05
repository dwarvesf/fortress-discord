package ogif

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (o *Ogif) RenderOgifLeaderboard(msg *model.DiscordMessage, leaderboard []model.OgifLeaderboardRecord, timeAmount int, timeUnit string) error {
	content := ""

	// Title
	title := "<:pepegod:819802587062468668> OGIF Leaderboard <:pepegod:819802587062468668>"

	// Leaderboard content
	if len(leaderboard) == 0 {
		content += fmt.Sprintf("No OGIF in the last %d %s.\n", timeAmount, timeUnit)
	} else {
		for i, ogif := range leaderboard {
			if i >= 20 {
				break
			}
			content += fmt.Sprintf("%d. <@%s> - %d OGIFs\n", i+1, ogif.DiscordID, ogif.SpeakCount)
		}

		if len(leaderboard) > 20 {
			content += "...\n"
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: content,
		Color:       0x00ff00, // Green color
	}

	return base.SendEmbededMessage(o.ses, msg, embed)
}
