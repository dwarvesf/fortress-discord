package project

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Project struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) ProjectViewer {
	return &Project{
		ses: ses,
	}
}

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
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Project Commands Help**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Project) List(original *model.DiscordMessage, projects []model.Project, page int) error {
	var content string
	pages := len(projects) / 10
	if len(projects)%10 != 0 {
		pages++
	}

	if len(projects) > 10 {
		projects = projects[(page-1)*10 : page*10]
	}

	for i := range projects {
		if projects[i].Name == "" {
			continue
		}

		p := projects[i]
		content += fmt.Sprintf("- %s", p.Name)

		if p.ArtifactLink != "" {
			content += fmt.Sprintf(" | [file](%s)", p.ArtifactLink)
		} else {
			content += " | file*"
		}

		if p.SourceLink != "" {
			content += fmt.Sprintf(" | [src](%s)", p.SourceLink)
		} else {
			content += " | src*"
		}

		if p.DocLink != "" {
			content += fmt.Sprintf(" | [doc](%s)", p.DocLink)
		} else {
			content += " | doc*"
		}

		content += "\n"
	}

	msg := &discordgo.MessageEmbed{
		Title:       "Projects",
		Description: content,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Page %d of %d", page, pages),
		},
	}

	base.SendEmbededMessage(e.ses, original, msg)
	return nil
}
