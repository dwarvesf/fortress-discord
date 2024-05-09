package event

import (
	"errors"
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Event) CreateGuildScheduledEvent(ev *model.DiscordEvent) error {
	e.l.Field("name", ev.Name).Info("create a scheduled event on discord")
	// get response from fortress
	err := e.adapter.Fortress().CreateGuildScheduledEvent(ev)
	if err != nil {
		e.l.Error(err, "can't create a scheduled event on discord")
		return err
	}

	return nil
}

func (e *Event) GetGuildScheduledEvents() ([]*model.DiscordEvent, error) {
	return e.adapter.Fortress().GetGuildScheduledEvents()
}

func (e *Event) SetSpeakers(message *model.DiscordMessage) error {
	errMsg := "Please provide speakers correctly! e.g `?event scheduled set speaker <discord_event_id> <@user1>:topic,<@user2>:topic_2`"
	if len(message.ContentArgs) < 6 {
		return errors.New(errMsg)
	}
	data := message.ContentArgs[5:]
	speakers := extractSpeakers(data)
	return e.adapter.Fortress().SetSpeakers(message.ContentArgs[4], speakers)
}

func extractSpeakers(data []string) []string {
	var result = make(map[string][]string, 0)
	for i, str := range data {
		if strings.HasPrefix(str, "<@") {
			// User ID found
			userId := str[2:strings.Index(str, ">")]
			// Topic found
			topic := data[i+2:]
			result[userId] = topic
		}
	}
	var speakers []string
	for u, v := range result {
		s := u + ":"
		for i := range v {
			if strings.Contains(v[i], "<@") {
				break
			}
			// remove elements from here
			s += v[i] + " "
		}
		speakers = append(speakers, s)
	}
	return speakers
}
