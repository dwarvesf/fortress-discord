package withdrawal

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Withdraw) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Withdrawal Commands**",
		"",
		"`?withdraw` - Initiate withdrawal process",
		"`?withdraw help` - Show this help message",
		"",
		"**Aliases**:",
		"`?withdraw h` - Shorthand for help",
		"",
		"**Requirements**:",
		"- Must have sufficient ICY token balance",
		"- Withdrawal amount depends on current ICY/VND exchange rate",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
