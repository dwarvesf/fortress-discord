package event

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Event) ListScheduledEvents(original *model.DiscordMessage, events []*model.Event) error {
	var content string

	for i := range events {
		t := ""
		if !events[i].Date.HasTime {
			t = events[i].Date.Time.Format("02 Jan 2006")
		} else {
			t = events[i].Date.Time.Format(time.RFC822)
		}

		content += fmt.Sprintf("%v - **%v** - %s \n", events[i].Id, events[i].Name, t)
	}

	if content == "" {
		content = "No scheduled events found"
	}

	msg := &discordgo.MessageEmbed{
		Title:       "Scheduled Events",
		Description: content,
		Color:       0x00ff00,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
