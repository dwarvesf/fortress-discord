package trend

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *Trend) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?trend <lang>** - Get daily trending Repositories in a particular programming language, in English.",
		"*Example:* `?trend go`\n",
		"**?trend <lang> <date_range>** - Get trending repository in English, filtered by programming language and date range(daily(default)/weekly/monthly)",
		"*Example:* `?trend go weekly`\n",
		"**?trend <lang> <date_range> <spoken_lang>** - Get trending repository, filtered by programming language, date range and spoken language(default: English)",
		"*Example*: `?trend go monthly en`\n",
		"**?trend programming** - Get available programming language",
		"**?trend spoken** - Get available spoken language",
		"**?trend date** - Get available date range",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}
	return base.SendEmbededMessage(t.ses, message, msg)
}
