package ai

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v view) Help(original *model.DiscordMessage) error {
	content := []string{
		"**AI Commands**",
		"",
		"`?ai <query>` - Ask an AI assistant a question or request help",
		"`?ai help` - Show this help message",
		"",
		"**Aliases**:",
		"`?ai h` - Shorthand for help",
		"",
		"**Example**:",
		"`?ai What is the capital of France?` - Ask an AI question",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
