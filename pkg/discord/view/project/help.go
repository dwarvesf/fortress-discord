package project

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Project Commands**",
		"",
		"`?project` - Show project commands help",
		"`?project list [status] [page]` - List projects (default: active status, page 1). Status: active, paused, closed",
		"`?project pnl` - Show project P&L information",
		"`?project commission` - Show project commission models",
		"",
		"**Aliases**",
		"`?project ls [status] [page]` - Alias for list (same arguments as list)",
		"`?project com` - Alias for commission",
		"`?project h` - Alias for help",
		"",
		"**Example**:",
		"`?project list active 1` - List active projects on the first page",
		"`?project pnl` - Show project profit and loss information",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Project) MilestoneHelp(message *model.DiscordMessage) error {
	content := []string{
		"**Milestone Commands**",
		"",
		"`?milestone list` - List project milestones",
		"`?milestone help` - Show this help message",
		"",
		"**Aliases**",
		"`?milestones` - Alias for milestone",
		"`?milestone ls` - Shorthand for list",
		"`?milestone h` - Shorthand for help",
		"",
		"**Permissions**:",
		"- Requires Mod or higher role",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
