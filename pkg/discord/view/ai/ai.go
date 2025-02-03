package ai

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type view struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) AIViewer {
	return &view{
		ses: ses,
	}
}

func (a *view) SendResponse(message *model.DiscordMessage, response *model.AIResponse) error {
	msg := &discordgo.MessageEmbed{
		Title:       "AI Response",
		Description: response.Response,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Input",
				Value:  response.Input,
				Inline: false,
			},
		},
	}

	return base.SendEmbededMessage(a.ses, message, msg)
}
