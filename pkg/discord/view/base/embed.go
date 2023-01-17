package base

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func SendEmbededMessage(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed) error {
	_, err := ses.ChannelMessageSendEmbed(original.ChannelId, normalize(original, embed))
	return err
}

// normalize add some default to embeded message if not set
func normalize(original *model.DiscordMessage, response *discordgo.MessageEmbed) *discordgo.MessageEmbed {
	if response.Timestamp == "" {
		response.Timestamp = time.Now().Format(time.RFC3339)
	}
	if response.Color == 0 {
		// default df color #D14960
		response.Color = 13715808
	}
	if response.Author == nil {
		response.Author = &discordgo.MessageEmbedAuthor{
			Name:    original.Author.Username,
			IconURL: original.Author.AvatarURL("128"),
		}
	}
	return response
}
