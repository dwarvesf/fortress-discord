package topic

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type TopicViewer interface {
	List(original *model.DiscordMessage, page, size string, researchTopic model.DiscordResearchTopicResponse) error
	BuildMessage(page, size string, researchTopic model.DiscordResearchTopicResponse) (msg *discordgo.MessageEmbed, components []discordgo.MessageComponent)
}
