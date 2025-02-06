package profile

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *Profile) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Profile Commands**",
		"",
		"`?profile @user` - Get employee profile",
		"`?profile help` - Show this help message",
		"",
		"**Permissions**:",
		"- Requires Smod or higher role in production",
		"- Limited to specific whitelisted channels",
		"",
		"**Aliases**:",
		"`?profile h` - Shorthand for help",
		"",
		"**Example**:",
		"`?profile @nam` - View profile of mentioned user",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}
