package mma

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *MMA) Help(message *model.DiscordMessage) error {
	content := []string{
		"**MMA Commands**",
		"",
		"`?mma template` - Export CSV template",
		"`?mma help` - Show this help message",
		"",
		"**Permissions**:",
		"- Requires Smod or higher role in production",
		"- Limited to specific whitelisted channels",
		"",
		"**Aliases**:",
		"`?mma h` - Shorthand for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}
