package ogif

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (o *Ogif) RenderOgifStats(original *model.DiscordMessage, userID string, stats model.OgifStats, timeAmount int, timeUnit string) error {
	content := ""
	// Overall stats
	content += "**Overall OGIF Stats**\n"
	lastPeriodStr := fmt.Sprintf("Last %d %s OGIFs", timeAmount, timeUnit)
	content += fmt.Sprintf("`Total OGIFs%*s`: %d\n", len(lastPeriodStr)-len("Total OGIFs"), "", stats.TotalSpeakCount)
	content += fmt.Sprintf("`%s`: %d\n\n", lastPeriodStr, stats.CurrentSpeakCount)

	// User stats
	userTag := "Your"
	if userID != original.Author.ID {
		userTag = fmt.Sprintf("<@%s>'s", userID)
	}
	content += fmt.Sprintf("**%s OGIF Stats**\n", userTag)
	allTimeRank := fmt.Sprintf("#%d", stats.UserAllTimeRank)
	if stats.UserAllTimeRank == 0 {
		allTimeRank = "No Rank"
	}
	content += fmt.Sprintf("`Total OGIFs%*s`: %d (Rank: %s)\n", len(lastPeriodStr)-len("Total OGIFs"), "", stats.UserAllTimeSpeaksCount, allTimeRank)

	currentRank := fmt.Sprintf("#%d", stats.UserCurrentRank)
	if stats.UserCurrentRank == 0 {
		currentRank = "No Rank"
	}
	content += fmt.Sprintf("`%s`: %d (Rank: %s)\n\n", lastPeriodStr, stats.UserCurrentSpeaksCount, currentRank)

	// OGIF list for the specified time period
	content += fmt.Sprintf("**%s OGIFs in the last %d %s**\n", userTag, timeAmount, timeUnit)
	if len(stats.OgifList) == 0 {
		content += fmt.Sprintf("No OGIFs found in the last %d %s.\n", timeAmount, timeUnit)
	} else {
		for i, speaker := range stats.OgifList {
			if i >= 20 {
				break
			}
			content += fmt.Sprintf("%d. %s\n", i+1, speaker.Topic)
		}

		if len(stats.OgifList) > 20 {
			content += "...\n"
		}
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepenote:885515949673951282> OGIF Stats <:pepenote:885515949673951282>",
		Description: content,
		Color:       0x00ff00, // Green color
	}

	return base.SendEmbededMessage(o.ses, original, msg)
}
