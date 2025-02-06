package event

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Event) Prefix() []string {
	return []string{"event", "events", "e"}
}

// Execute is where we handle logic for each command
func (e *Event) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?event`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?event list`
	switch message.ContentArgs[1] {
	// TODO: notion upcoming events, check if it's still needed
	case "upcoming":
		return e.List(message)
	case "list", "ls":
		return e.ListGuildScheduledEvents(message)
	case "help", "h":
		return e.Help(message)
	case "speakerset", "spks":
		return e.SetSpeakers(message)
	default:
		return e.view.Done().MissingContent(message)
	}
}

func (e *Event) Name() string {
	return "Event Command"
}

func (e *Event) Help(message *model.DiscordMessage) error {
	return e.view.Event().Help(message)
}

func (e *Event) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *Event) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
