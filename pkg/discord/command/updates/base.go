package updates

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (u *Updates) Prefix() []string {
	return []string{"update", "updates"}
}

// Execute is where we handle logic for each command
func (u *Updates) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return u.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return u.List(message)
	}

	return nil
}

func (u *Updates) Name() string {
	return "Updates Command"
}

func (u *Updates) Help(message *model.DiscordMessage) error {
	return nil
}

func (u *Updates) DefaultCommand(message *model.DiscordMessage) error {
	return u.List(message)
}

func (u *Updates) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
