package event

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Event) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Event Commands**",
		"",
		"`?event list` - List Discord scheduled events",
		"`?event speakerset` - Set speakers for scheduled events",
		"`?event help` - Show this help message",
		"",
		"**Aliases**:",
		"`?events` - Alternative for event command",
		"`?e` - Alternative for event command",
		"`?events ls` - Shorthand for list",
		"`?e spks` - Shorthand for speakerset",
		"`?event h` - Shorthand for help",
		"",
		"**Speaker Set Command Format**:",
		"`?event speakerset <eventID> <topic1> @user1 @user2 <topic2> @user3`",
		"",
		"**Example**:",
		"`?event list` - View scheduled events",
		"`?event spks 1256184313876054147 Erlang @bienvh AI-101 @cherry` - Set speakers for an event",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
