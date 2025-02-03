package issue

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (i *Issue) Prefix() []string {
	return []string{"issue", "issues"}
}

// Execute is where we handle logic for each command
func (i *Issue) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return i.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list", "ls":
		return i.List(message)
	case "help", "h":
		return i.Help(message)
	default:
		return i.Help(message)
	}
}

func (i *Issue) Name() string {
	return "Issue Command"
}

func (i *Issue) Help(message *model.DiscordMessage) error {
	return i.view.Issue().Help(message)
}

func (i *Issue) DefaultCommand(message *model.DiscordMessage) error {
	return i.Help(message)
}

func (i *Issue) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	// we require roles for all commands in milestones
	return permutil.CheckSupporterOrAbove(message.Roles)
}
