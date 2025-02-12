package df

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type view struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) DFViewer {
	return &view{
		ses: ses,
	}
}

func (a *view) SendResponse(message *model.DiscordMessage, response *model.AIResponse) error {
	msg := &discordgo.MessageEmbed{
		Title:       "DF Assistant Response",
		Description: response.Response,
	}

	return base.SendEmbededMessage(a.ses, message, msg)
}
