package hiring

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Hiring) Prefix() []string {
	return []string{"hiring", "jobs"}
}

// Execute is where we handle logic for each command
func (e *Hiring) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?hiring`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?hiring list`
	switch message.ContentArgs[1] {
	case "list", "ls":
		return e.OpenPositions(message)
	case "help", "h":
		return e.Help(message)
	default:
		return e.Help(message)
	}
}

func (e *Hiring) Name() string {
	return "Hiring Command"
}

func (e *Hiring) Help(message *model.DiscordMessage) error {
	return nil
}

func (e *Hiring) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *Hiring) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
