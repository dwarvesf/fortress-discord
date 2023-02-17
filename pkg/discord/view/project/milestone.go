package project

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) ListMilestones(original *model.DiscordMessage, milestones []*model.ProjectMilestone) error {
	for i := range milestones {
		v := milestones[i]
		if len(v.Milestones) == 0 {
			continue
		}
		var content string
		for _, vv := range v.Milestones {
			t := ""
			past := false
			if vv.EndDate != nil {
				t = vv.EndDate.Format("02 Jan 2006")
				if vv.EndDate.Before(time.Now()) {
					past = true
				} else {
					past = false
				}
			}

			content += fmt.Sprintf("[%s]・ **%s**", t, vv.Name)
			if past {
				content += "<:pepeno2:885513214467661834>"
			}
			content += "\n"

			for _, subMilestone := range vv.SubMilestones {
				content += fmt.Sprintf("\t- %s\n", subMilestone.Name)
			}
		}
		msg := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("<:pepe_ping:1028964391690965012> %s <:pepe_ping:1028964391690965012>", v.Name),
			Description: content,
		}
		base.SendEmbededMessage(e.ses, original, msg)
	}

	return nil
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
