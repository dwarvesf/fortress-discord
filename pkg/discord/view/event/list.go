package event

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Event) List(original *model.DiscordMessage, events []*model.Event) error {
	var content string

	for i := range events {
		t := ""
		if !events[i].Date.HasTime {
			t = events[i].Date.Time.Format("02 Jan 2006")
		} else {
			t = events[i].Date.Time.Format(time.RFC822)
		}

		content += fmt.Sprintf("**%s** ・ %s \n", events[i].Name, t)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepe_ping:1028964391690965012> Upcoming Events <:pepe_ping:1028964391690965012>",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
