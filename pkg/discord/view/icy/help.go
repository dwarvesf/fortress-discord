package icy

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Icy) Help(message *model.DiscordMessage) error {
	content := []string{
		"**ICY Token Commands**",
		"",
		"`?icy list` - List ICY token transactions",
		"`?icy info` - View personal ICY token information",
		"`?icy help` - Show this help message",
		"",
		"**Restricted Commands**:",
		"`?icy accounting` - View ICY token accounting (Smod+ only)",
		"",
		"**Aliases**:",
		"`?icy ls` - Shorthand for list",
		"`?icy i` - Shorthand for info",
		"`?icy h` - Shorthand for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
