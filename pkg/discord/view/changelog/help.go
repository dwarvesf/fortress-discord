package changelog

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c *Changelog) Help(original *model.DiscordMessage) error {
	content := []string{
		"**Changelog Commands**",
		"",
		"`?changelog send` - Send changelog data, render, filter input, process, and send",
		"`?changelog help` - Show this help message",
		"",
		"**Aliases**:",
		"`?changelog h` - Shorthand for help",
		"",
		"**Example**:",
		"`?changelog` - View list of recent changes",
		"`1,3,5` - Send changelogs with numbers 1, 3, and 5",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(c.ses, original, msg)
}
