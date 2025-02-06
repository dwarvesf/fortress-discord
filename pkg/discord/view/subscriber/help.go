package subscriber

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (s *Subscriber) Help(original *model.DiscordMessage) error {
	content := []string{
		"**Subscriber Commands**",
		"",
		"`?new list` - List new subscribers",
		"`?new help` - Show this help message",
		"",
		"**Aliases**:",
		"`?new ls` - Shorthand for list",
		"`?new h` - Shorthand for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
