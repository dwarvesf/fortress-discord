package salary

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (s *Salary) Help(original *model.DiscordMessage) error {
	content := []string{
		"**Salary Advance Commands**",
		"",
		"`?salary advance` - Request a salary advance",
		"`?salary help` - Show this help message",
		"",
		"**Aliases**:",
		"`?salary adv` - Shorthand for salary advance",
		"`?salary h` - Shorthand for help",
		"",
		"**Example**:",
		"`?salary advance` - Initiate a salary advance request",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
