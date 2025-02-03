package milestone

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (m *Milestone) Prefix() []string {
	return []string{"milestone", "milestones"}
}

// Execute is where we handle logic for each command
func (m *Milestone) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return m.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "list", "ls":
		return m.ListMilestones(message)
	case "help", "h":
		return m.Help(message)
	default:
		return m.DefaultCommand(message)
	}
}

func (m *Milestone) Name() string {
	return "Milestone Command"
}

func (m *Milestone) Help(message *model.DiscordMessage) error {
	return nil
}

func (m *Milestone) DefaultCommand(message *model.DiscordMessage) error {
	return m.Help(message)
}

func (m *Milestone) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	// we require roles for all commands in milestones
	return permutil.CheckModOrAbove(message.Roles)
}
