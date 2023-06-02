package icy

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Icy) List(original *model.DiscordMessage, icys []*model.Icy) error {
	var content string

	if len(icys) > 0 {
		date, err := time.Parse("2006-01-02", icys[0].Period)
		if err != nil {
			return err
		}
		plusSevenDays := date.AddDate(0, 0, 7)
		endDate := plusSevenDays.Format("2006 January 2")
		startDate := date.Format("2006 January 2")
		content += fmt.Sprintf("\n**%s** - **%s**\n\n", startDate, endDate)
	}

	for i := range icys {
		icy := icys[i]

		content += fmt.Sprintf("・ %s Team: %s :ice_cube:\n", strings.Title(icy.Team), icy.Amount)
	}

	content += "\nHead to [earn.d.foundation](https://earn.d.foundation) to see available quests and r&d topics.\n"

	msg := &discordgo.MessageEmbed{
		Title:       "This week ICY :ice_cube:",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
