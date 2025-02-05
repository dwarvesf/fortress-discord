package sum

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Sum) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Article Summarization Commands**",
		"",
		"`?sum <article_url>` - Summarize an article from a given URL",
		"`?sum help` - Show this help message",
		"",
		"**Aliases**:",
		"`?sum h` - Shorthand for help",
		"",
		"**Example**:",
		"`?sum https://example.com/article` - Summarize the article at the specified URL",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
