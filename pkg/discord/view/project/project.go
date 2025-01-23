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
		"`?project list` - List all projects",
		"`?project pnl` - Show project P&L information",
		"`?project commission` - Show project commission models",
		"",
		"**Aliases**",
		"`?project l` - Alias for list",
		"`?project com` - Alias for commission",
		"`?project h` - Alias for help",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Project Commands Help**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Project) List(original *model.DiscordMessage, projects []model.Project) error {
	var content string
	for i := range projects {
		p := projects[i]
		content += fmt.Sprintf("`Name.        ` **%s**\n", p.Name)
		content += fmt.Sprintf("`Code.        ` **%s**\n", p.Code)
		if p.ArtifactLink != "" {
			content += fmt.Sprintf("`Artifacts.   ` [Link](%s)\n", p.ArtifactLink)
		}
		content += "\n"
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:pepe_ping:1028964391690965012> Projects <:pepe_ping:1028964391690965012>"),
		Description: content,
	}

	base.SendEmbededMessage(e.ses, original, msg)
	return nil
}
