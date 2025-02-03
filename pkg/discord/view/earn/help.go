package earn

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Earn) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Earn Commands**",
		"",
		"`?earn list` - List available earning opportunities",
		"`?earn help` - Show this help message",
		"",
		"**Aliases**:",
		"`?earn ls` - Shorthand for list",
		"`?earn h` - Shorthand for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
