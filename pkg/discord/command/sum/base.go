package sum

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Sum) Prefix() []string {
	return []string{"sum"}
}

// Execute is where we handle logic for each command
func (e *Sum) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?sum`
	if len(message.ContentArgs) == 1 {
		return e.Help(message)
	}

	// handle command for 2 args input from user, e.g `?earn sum`
	switch message.ContentArgs[1] {
	case "help", "h":
		return e.Help(message)
	}

	return e.Sum(message)
}

func (e *Sum) Name() string {
	return "Sum Command"
}

func (e *Sum) Help(message *model.DiscordMessage) error {
	return e.view.Sum().Help(message)
}

func (e *Sum) DefaultCommand(message *model.DiscordMessage) error {
	return e.Sum(message)
}

func (e *Sum) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
