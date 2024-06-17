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

	// upcoming event
	content += "**Upcoming Events**\n"

	upcomingEvents := make([]model.Event, 0)
	for i := range events {
		if events[i] != nil && !events[i].IsOver {
			upcomingEvents = append(upcomingEvents, *events[i])
		}
	}

	content += listEventsContent(upcomingEvents)

	// Passed event
	content += "\n**Passed Events**\n"

	passedEvents := make([]model.Event, 0)
	for i := range events {
		if events[i] != nil && events[i].IsOver {
			passedEvents = append(passedEvents, *events[i])
		}
	}

	content += listEventsContent(passedEvents)
	content += "\n"

	msg := &discordgo.MessageEmbed{
		Title:       "Scheduled Events",
		Description: content,
		Color:       0x00ff00,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func listEventsContent(events []model.Event) string {
	content := ""

	for i := range events {
		t := ""
		if !events[i].Date.HasTime {
			t = events[i].Date.Time.Format("02 Jan 2006")
		} else {
			t = events[i].Date.Time.Format(time.RFC822)
		}

		content += fmt.Sprintf("🔹 `%v` **%v**\n", events[i].Id, events[i].Name)
		content += fmt.Sprintf("\u2570  `%s`\n", t)
	}

	if len(events) == 0 {
		content += "No events found"
	}

	return content
}
