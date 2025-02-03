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
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}

func (t *TechRadar) AssessHelp(original *model.DiscordMessage) error {
	content := []string{
		"**Assess Quadrant Commands**",
		"",
		"`?assess list` - List technologies in the Assess quadrant",
		"`?assess help` - Show this help message",
		"",
		"**Aliases**:",
		"`?assess ls` - Alias for assess list",
		"`?assess h` - Alias for assess help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}

func (t *TechRadar) HoldHelp(original *model.DiscordMessage) error {
	content := []string{
		"**Hold Quadrant Commands**",
		"",
		"`?hold list` - List technologies in the Hold quadrant",
		"`?hold help` - Show this help message",
		"",
		"**Aliases**:",
		"`?hold ls` - Alias for hold list",
		"`?hold h` - Alias for hold help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}

func (t *TechRadar) TrialHelp(original *model.DiscordMessage) error {
	content := []string{
		"**Trial Quadrant Commands**",
		"",
		"`?trial list` - List technologies in the Trial quadrant",
		"`?trial help` - Show this help message",
		"",
		"**Aliases**:",
		"`?trial ls` - Alias for trial list",
		"`?trial h` - Alias for trial help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}

func (t *TechRadar) IndexHelp(original *model.DiscordMessage) error {
	content := []string{
		"**Index Commands**",
		"",
		"`?index list` - List topics in the Tech Radar",
		"`?index help` - Show this help message",
		"",
		"**Aliases**:",
		"`?index ls` - Alias for index list",
		"`?index h` - Alias for index help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}
