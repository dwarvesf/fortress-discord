package news

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v view) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?news <platform> <topic>** - get news from given platform with specific topic",
		"",
		"*Example*: `?news reddit golang`",
		"",
		"Currently supported platforms:",
		"- [Reddit](https://www.reddit.com/)",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}
