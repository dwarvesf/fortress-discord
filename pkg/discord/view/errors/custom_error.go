package errors

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Error) Raise(original *model.DiscordMessage, errorMessage string) error {
	msg := &discordgo.MessageEmbed{
		Title:       ":warning: **Error Message** :warning: ",
		Description: errorMessage,
	}

	return base.SendEmbededMessage(e.ses, original, msg)

}
