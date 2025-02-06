package topic

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type TopicViewer interface {
	List(original *model.DiscordMessage, timeRange string, researchTopic model.DiscordResearchTopicResponse) error
	BuildMessage(timeRange string, researchTopic model.DiscordResearchTopicResponse) (msg *discordgo.MessageEmbed, components []discordgo.MessageComponent)
	Help(original *model.DiscordMessage) error
}
