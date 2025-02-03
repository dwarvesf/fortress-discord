package done

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (d *Done) Prefix() []string {
	return []string{"done"}
}

// Execute is where we handle logic for each command
func (d *Done) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return d.Help(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "help", "h":
		return d.Help(message)
	default:
		return d.Done(message)
	}
}

func (d *Done) Name() string {
	return "Done Command"
}

func (d *Done) Help(message *model.DiscordMessage) error {
	return d.view.Done().Help(message)
}

func (d *Done) DefaultCommand(message *model.DiscordMessage) error {
	return d.Help(message)
}

func (d *Done) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	// we require roles for all commands in milestones
	return permutil.CheckSupporterOrAbove(message.Roles)
}
