package icy

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Icy) List(original *model.DiscordMessage, icys []*model.Icy) error {
	var content string
	for i := range icys {
		icy := icys[i]

		date, err := time.Parse("2006-01-02", icy.Period)
		if err != nil {
			return err
		}
		plusSevenDays := date.AddDate(0, 0, 7)
		endDate := plusSevenDays.Format("2006/01/02")
		startDate := date.Format("2006/01/02")

		content += fmt.Sprintf("**%s** :ice_cube: are distributed to **%s** (**%s** - **%s**)\n", icy.Amount, icy.Team, startDate, endDate)
	}

	msg := &discordgo.MessageEmbed{
		Title:       ":ice_cube: **Weekly Icy Distribution** :ice_cube: ",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
