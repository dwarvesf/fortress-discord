package base

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func SendEmbededMessage(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed) error {
	return SendEmbededMessageWithChannel(ses, original, embed, original.ChannelId)
}

func SendComplexMessage(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed, components []discordgo.MessageComponent) error {
	return SendComplexMessageWithChannel(ses, original, embed, components, original.ChannelId)
}

func SendComplexMessageWithChannel(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed, components []discordgo.MessageComponent, channelId string) error {
	_, err := ses.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
		Embed:      Normalize(ses, embed),
		Components: components,
	})
	return err
}

func SendEmbededMessageWithChannel(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed, channelId string) error {
	_, err := ses.ChannelMessageSendEmbed(channelId, Normalize(ses, embed))
	return err
}

func SendComplexReplyMessageWithChannel(ses *discordgo.Session, original *model.DiscordMessage, embed *discordgo.MessageEmbed, components []discordgo.MessageComponent, channelId string) error {
	msgData := &discordgo.MessageSend{
		Embed:      Normalize(ses, embed),
		Components: components,
		Reference: &discordgo.MessageReference{
			MessageID: original.MessageId,
			ChannelID: channelId,
			GuildID:   original.GuildId,
		},
	}

	_, err := ses.ChannelMessageSendComplex(channelId, msgData)
	return err
}

func SendMessage(ses *discordgo.Session, original *model.DiscordMessage, msg *discordgo.Message) error {
	return SendMessageWithChannel(ses, original, msg, original.ChannelId)
}

func SendMessageWithChannel(ses *discordgo.Session, original *model.DiscordMessage, msg *discordgo.Message, channelId string) error {
	_, err := ses.ChannelMessageSend(channelId, msg.Content)
	return err
}

// Normalize add some default to embeded message if not set
func Normalize(s *discordgo.Session, response *discordgo.MessageEmbed) *discordgo.MessageEmbed {
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
			IconURL: s.State.User.AvatarURL("128"),
			Text:    "?help to see all commands",
		}
	}
	return response
}
