package ogif

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Ogif) Help(message *model.DiscordMessage) error {
	embed := &discordgo.MessageEmbed{
		Title:       "OGIF Command Help",
		Description: "The OGIF command allows you to fetch and display OGIF stats for a user or view the leaderboard.",
		Color:       0x00ff00, // Green color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Usage",
				Value: "`ogif [user_mention|top] [time_period]`",
			},
			{
				Name:  "Parameters",
				Value: "- `user_mention`: Optional. Mention the user to fetch stats for. If omitted, uses the command author.\n- `top`: Optional. Use this to view the leaderboard instead of user stats.\n- `time_period`: Optional. Time period for stats (e.g., '7d', '30d', '3m'). Default is 30 days.",
			},
			{
				Name:  "Examples",
				Value: "- `ogif`\n- `ogif @user`\n- `ogif @user 7d`\n- `ogif 14d`\n- `ogif top`\n- `ogif top 7d`",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "For more information, contact the bot administrator.",
		},
	}

	return base.SendEmbededMessage(e.ses, message, embed)
}
