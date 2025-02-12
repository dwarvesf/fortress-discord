package df

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v view) Help(original *model.DiscordMessage) error {
	content := []string{
		"**DF Commands**",
		"",
		"`?df <query>` - Ask an DF employee assistant a question",
		"`?df help` - Show this help message",
		"",
		"**Aliases**:",
		"`?df h` - Shorthand for help",
		"",
		"**Example**:",
		"`?df How many employees taking day-off today?` - Ask an employee question",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
