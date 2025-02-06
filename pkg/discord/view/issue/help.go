package issue

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (i *Issue) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Issue Commands**",
		"",
		"`?issue list` - List current issues",
		"`?issue help` - Show this help message",
		"",
		"**Aliases**:",
		"`?issues list` - Alternative for issue list",
		"`?issue ls` - Shorthand for list",
		"`?issue h` - Shorthand for help",
		"",
		"**Permissions**:",
		"- Requires Supporter or higher role",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(i.ses, message, msg)
}
