package techradar

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *TechRadar) AdoptHelp(original *model.DiscordMessage) error {
	content := []string{
		"**Adopt Quadrant Commands**",
		"",
		"`?adopt list` - List technologies in the Adopt quadrant",
		"`?adopt help` - Show this help message",
		"",
		"**Aliases**:",
		"`?adopt ls` - Alias for adopt list",
		"`?adopt h` - Alias for adopt help",
		"",
		"**Example**:",
		"`?adopt list` - List technologies in the Adopt quadrant",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}
