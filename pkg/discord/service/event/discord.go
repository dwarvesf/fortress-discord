package event

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
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

// SetSpeakers is a command to set speakers for a scheduled event
// Input: `?event speakerset <eventID>  <topic1> @user1 @user2 <topic2> @user3`
func (e *Event) SetSpeakers(message *model.DiscordMessage) error {
	errMsg := "Please provide speakers correctly! e.g `event speakerset <eventID> <topic1> @user1 @user2 <topic2> @user3`"

	if len(message.ContentArgs) < 5 {
		return errors.New(errMsg)
	}

	mapSpeakers := extractSpeakers(message.ContentArgs)
	event, err := e.adapter.Fortress().SetSpeakers(message.ContentArgs[2], mapSpeakers)
	if err != nil {
		e.l.Error(err, "can't set speakers for the scheduled event")
		return err
	}

	content := fmt.Sprintf("`Event:       ` **%s**\n", event.Name)
	content += fmt.Sprintf("`Description: ` **%s**\n", event.Description)

	if len(mapSpeakers) != 0 {
		content += "\n"
		content += "**Topics**\n"
	}

	for topic, speakers := range mapSpeakers {
		content += fmt.Sprintf("- **%s**: ", topic)
		for _, speaker := range speakers {
			content += "<@" + speaker + ">"
		}
		content += "\n"
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:check:1077631110047080478> Set Speakers Successfully",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func extractSpeakers(args []string) map[string][]string {
	// ignore 3 first elements of data that are command and event id `event speakerset <event_id>`
	if len(args) <= 3 {
		return nil
	}

	args = args[3:]

	// Normalize input by removing all empty
	var normalizedData []string
	for _, d := range args {
		if d != "" {
			normalizedData = append(normalizedData, d)
		}
	}

	// Extract speakers
	mapSpeakersByTopic := make(map[string][]string)
	currentTopic := ""
	isNewTopic := true
	for _, str := range normalizedData {
		if !strings.Contains(str, "<@") {
			if isNewTopic {
				currentTopic = str
				isNewTopic = false
				continue
			}
			currentTopic += " " + str
			continue
		}

		if currentTopic == "" {
			continue
		}

		isNewTopic = true
		mapSpeakersByTopic[currentTopic] = append(mapSpeakersByTopic[currentTopic], strings.Trim(str, "<@>"))
	}

	return mapSpeakersByTopic
}

func (e *Event) GetOgifStats(discordID string, after time.Time) (model.OgifStats, error) {
	resp, err := e.adapter.Fortress().GetOgifStats(discordID, after)
	if err != nil {
		return model.OgifStats{}, err
	}
	return resp.Data, err
}

func (e *Event) GetOgifLeaderboard(after time.Time, limit int) ([]model.OgifLeaderboardRecord, error) {
	leaderboard, err := e.adapter.Fortress().GetOgifLeaderboard(after, limit)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}
