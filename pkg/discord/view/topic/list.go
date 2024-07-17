package topic

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *Topic) BuildMessage(page, size string, researchTopic model.DiscordResearchTopicResponse) (msg *discordgo.MessageEmbed, components []discordgo.MessageComponent) {
	var content string

	content += "These are most discussed research topic in Dwarves Foundation in the last 7 days.\n\n"

	topicsPerPage, _ := strconv.Atoi(size)
	pageInt, _ := strconv.Atoi(page)
	totalTopic := researchTopic.Total
	totalPages := (totalTopic + topicsPerPage - 1) / topicsPerPage
	topics := researchTopic.Data

	if pageInt < 1 {
		pageInt = 1
	} else if pageInt > totalPages {
		pageInt = totalPages
	}

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

	components = []discordgo.MessageComponent{}

	if totalTopic == 0 {
		msg.Description = "No active research topics found in the last 7 days."
		return msg, components
	}

	// Create pagination buttons
	if totalTopic > topicsPerPage {
		components = []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Previous",
						Style:    discordgo.PrimaryButton,
						CustomID: fmt.Sprintf("topic_prev_%d", pageInt),
						Disabled: pageInt == 1,
					},
					discordgo.Button{
						Label:    "Next",
						Style:    discordgo.PrimaryButton,
						CustomID: fmt.Sprintf("topic_next_%d", pageInt),
						Disabled: pageInt == totalPages,
					},
				},
			},
		}
	}

	return msg, components
}

func (v *Topic) List(original *model.DiscordMessage, page, size string, researchTopic model.DiscordResearchTopicResponse) error {
	msg, components := v.BuildMessage(page, size, researchTopic)

	return base.SendComplexMessage(v.ses, original, msg, components)
}
