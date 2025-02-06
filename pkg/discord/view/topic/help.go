package topic

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Topic) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Research Topic Commands**",
		"",
		"`?topic list` - Show the top research topics for the last 7 days",
		"`?topic help` - Show this help message",
		"",
		"**Aliases**:",
		"`?topics list` - Alternative for topic list",
		"`?topic l` - Shorthand for list",
		"`?topic h` - Shorthand for help",
		"",
		"**Example**:",
		"`?topic list` - View top research topics for the past week",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
