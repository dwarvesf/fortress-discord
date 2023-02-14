package issue

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (i *Issue) Prefix() []string {
	return []string{"issue"}
}

// Execute is where we handle logic for each command
func (i *Issue) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return i.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return i.List(message)
	}

	return nil
}

func (i *Issue) Name() string {
	return "Issue Command"
}

func (i *Issue) Help(message *model.DiscordMessage) error {
	return nil
}

func (i *Issue) DefaultCommand(message *model.DiscordMessage) error {
	return i.List(message)
}

func (i *Issue) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
