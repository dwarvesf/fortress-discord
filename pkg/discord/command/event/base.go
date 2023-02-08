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
	}

	return nil
}

func (e *Event) Name() string {
	return "Event Command"
}

func (e *Event) Help(message *model.DiscordMessage) error {
	return nil
}

func (e *Event) DefaultCommand(message *model.DiscordMessage) error {
	return e.List(message)
}

func (e *Event) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
