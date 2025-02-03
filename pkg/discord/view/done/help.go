package done

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (d *Done) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Done Commands**",
		"",
		"`?done <description>` - Mark a task or milestone as completed",
		"`?done help` - Show this help message",
		"**Aliases**:",
		"`?done h` - Shorthand for help",
		"",
		"**Example**:",
		"`?done Completed backend API refactoring` - Mark task as done",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(d.ses, message, msg)
}
