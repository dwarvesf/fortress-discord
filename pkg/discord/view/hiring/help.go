package hiring

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Hiring) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Hiring Commands**",
		"",
		"`?hiring list` - List open job positions",
		"`?hiring help` - Show this help message",
		"",
		"**Aliases**:",
		"`?jobs list` - Alternative for hiring list",
		"`?hiring ls` - Shorthand for list",
		"`?hiring h` - Shorthand for help",
		"",
		"**Example**:",
		"`?hiring list` - View current open job positions",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
