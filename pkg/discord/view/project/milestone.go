package project

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) ListMilestones(original *model.DiscordMessage, milestones *model.ProjectMilestone) error {
	var content string

	for i, v := range milestones.Milestones {
		t := ""
		if v.EndDate != nil {
			t = v.EndDate.Format("02 Jan 2006")
		}

		content += fmt.Sprintf("%d. **%s** ・ %s \n", i+1, milestones.Milestones[i].Name, t)

		for _, subMilestone := range v.SubMilestones {
			content += fmt.Sprintf("\t- %s\n", subMilestone.Name)
		}

	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepe_ping:1028964391690965012> Upcoming Milestones <:pepe_ping:1028964391690965012>",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func (e *Project) EmptyMilestones(original *model.DiscordMessage) error {
	content := "No milestones found, please make sure you have defined milestones for this project."

	msg := &discordgo.MessageEmbed{
		Title:       ":warning: **No Milestones Found** :warning:",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func (e *Project) MissingArgsMilestones(original *model.DiscordMessage) error {
	content := `No project name provided, please make sure you have provided a project name.

	e.g: **?milestones nghenhan**`

	msg := &discordgo.MessageEmbed{
		Title:       ":warning: **Missing Project Name** :warning:",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
