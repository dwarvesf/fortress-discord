package base

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func SendEmbededMessage(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed) error {
	return SendEmbededMessageWithChannel(ses, original, embed, original.ChannelId)
}

func SendEmbededMessageWithChannel(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed, channelId string) error {
	_, err := ses.ChannelMessageSendEmbed(channelId, Normalize(original.Author, embed))
	return err
}

// Normalize add some default to embeded message if not set
func Normalize(user *discordgo.User, response *discordgo.MessageEmbed) *discordgo.MessageEmbed {
	if response.Timestamp == "" {
		response.Timestamp = time.Now().Format(time.RFC3339)
	}

	// I did something tricky here, if timestamp is custom, we don't want to show it, because in case of user want to add a custom date time format in the footer
	// instead of automatically add it, we don't want to show it twice.
	if response.Timestamp == "custom" {
		response.Timestamp = ""
	}

	if response.Color == 0 {
		// default df color #D14960
		response.Color = 13715808
	}
	if response.Footer == nil {
		response.Footer = &discordgo.MessageEmbedFooter{
			IconURL: user.AvatarURL("128"),
			Text:    "?help to see all commands",
		}
	}
	return response
}
