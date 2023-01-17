package errors

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Error) CommandNotFound(original *model.DiscordMessage) error {
	content := "Please type `?help` to see all available commands."
	msg := &discordgo.MessageEmbed{
		Title:       ":warning: **Command not found** :warning: ",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)

}
