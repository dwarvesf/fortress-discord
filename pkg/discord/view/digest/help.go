package digest

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Digest) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Digest Commands**",
		"",
		"`?digest list` - List internal updates",
		"`?digest help` - Show this help message",
		"",
		"**Aliases**:",
		"`?digest ls` - Shorthand for digest list",
		"`?digest h` - Shorthand for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(a.ses, message, msg)
}
