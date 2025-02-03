package staff

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (s *Staff) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Staff Commands**",
		"",
		"`?staff list` - List staff members",
		"`?staff help` - Show this help message",
		"",
		"**Aliases**:",
		"`?staff ls` - Shorthand for list",
		"`?staff h` - Shorthand for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(s.ses, message, msg)
}
