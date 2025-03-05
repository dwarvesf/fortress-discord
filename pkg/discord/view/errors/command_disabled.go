package errors

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Error) CommandTemporarilyDisabled(original *model.DiscordMessage) error {
	content := "This command is temporarily disabled."
	msg := &discordgo.MessageEmbed{
		Title:       ":warning: **Command Temporarily Disabled** :warning:",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
