package base

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func SendEmbededMessage(ses *discordgo.Session, channelId string, embed *discordgo.MessageEmbed) error {
	_, err := ses.ChannelMessageSendEmbed(channelId, normalize(embed))
	return err
}

func normalize(msg *discordgo.MessageEmbed) *discordgo.MessageEmbed {
	if msg.Timestamp == "" {
		msg.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	}
	if msg.Color == 0 {
		// default df color #D14960
		msg.Color = 13715808
	}
	return msg
}
