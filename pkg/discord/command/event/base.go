package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Event) Prefix() []string {
	return []string{"event", "events"}
}

// Execute is where we handle logic for each command
func (e *Event) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?event`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?event list`
	switch message.ContentArgs[1] {
	case "list":
		return e.List(message)
	case "schedule", "scheduled":
		// handle command for 3 args input from user, e.g `?event scheduled list` by default
		if len(message.ContentArgs) == 2 {
			return e.ListGuildScheduledEvents(message)
		}
		switch message.ContentArgs[2] {
		case "list", "ls":
			return e.ListGuildScheduledEvents(message)
		case "set", "s":
			// default will set speaker, separate by space for multiple speakers, format: <discord_event_id> <@speaker_discord_id>:topic
			switch message.ContentArgs[3] {
			case "speaker", "spk":
				return e.SetSpeakers(message)
			default:
				return e.view.Done().MissingContent(message)
			}
		}
	case "help", "h":
		return e.Help(message)
	}

	return nil
}

func (e *Event) Name() string {
	return "Event Command"
}

func (e *Event) Help(message *model.DiscordMessage) error {
	return e.view.Event().Help(message)
}

func (e *Event) DefaultCommand(message *model.DiscordMessage) error {
	return e.List(message)
}

func (e *Event) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
