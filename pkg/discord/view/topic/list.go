package topic

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *Topic) BuildMessage(timeRange string, researchTopic model.DiscordResearchTopicResponse) (msg *discordgo.MessageEmbed, components []discordgo.MessageComponent) {
	var content string
	now := time.Now()

	description := fmt.Sprintf("These are the most discussed research topic in Dwarves Foundation in the last %s days.\n\n", timeRange)
	if timeRange == constant.AllTime {
		description = "These are the most discussed research topics in Dwarves Foundation of all time.\n\n"
	}
	content += description

	topics := researchTopic.Data

	for _, topic := range topics {
		content += fmt.Sprintf("%s - %d messages\n", topic.URL, topic.MsgCount)
		for _, user := range topic.SortedActiveUsers {
			authorsStr := "**@unknown-user**"
			if user.UserID != "" {
				authorsStr = fmt.Sprintf("<@%s>", user.UserID)
			}
			content += fmt.Sprintf("∟ %s - %d messages\n", authorsStr, user.MsgCount)
		}
		content += "\n"
	}

	msg = &discordgo.MessageEmbed{
		Title:       "<:uongcafe:819507964003876906> Active Topic Discussed \n",
		Description: content,
	}

	if len(topics) == 0 {
		msgDescription := fmt.Sprintf("There was no active topic in last %s days", timeRange)
		if timeRange == constant.AllTime {
			msgDescription = "There was no active topic at all times"
		}
		msg.Description = msgDescription
	}

	mapTimeRangeOption := map[string]string{
		"7":   "7 days",
		"30":  "30 days",
		"all": "All time",
	}

	selectTimeMenu := discordgo.SelectMenu{
		CustomID: "time_select",
		Options: []discordgo.SelectMenuOption{
			{
				Label:       "7 days",
				Value:       "7",
				Description: fmt.Sprintf("%s - %s", now.AddDate(0, 0, -7).Format("Jan 2, 2006"), now.Format("Jan 2, 2006")),
				Emoji:       &discordgo.ComponentEmoji{Name: "📅"},
			},
			{
				Label:       "30 days",
				Value:       "30",
				Description: fmt.Sprintf("%s - %s", now.AddDate(0, 0, -30).Format("Jan 2, 2006"), now.Format("Jan 2, 2006")),
				Emoji:       &discordgo.ComponentEmoji{Name: "📅"},
			},
			{
				Label:       "All time",
				Value:       "all",
				Description: "All time",
				Emoji:       &discordgo.ComponentEmoji{Name: "📅"},
			},
		},
		Placeholder: fmt.Sprintf("📅 %s", mapTimeRangeOption[timeRange]),
	}

	components = []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				selectTimeMenu,
			},
		},
	}

	return msg, components
}

func (v *Topic) List(original *model.DiscordMessage, timeRange string, researchTopic model.DiscordResearchTopicResponse) error {
	msg, components := v.BuildMessage(timeRange, researchTopic)

	return base.SendComplexMessage(v.ses, original, msg, components)
}
