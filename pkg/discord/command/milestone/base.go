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
	return m.ListMilestones(message)
}

func (m *Milestone) Name() string {
	return "Milestone Command"
}

func (m *Milestone) Help(message *model.DiscordMessage) error {
	return nil
}

func (m *Milestone) DefaultCommand(message *model.DiscordMessage) error {
	return m.ListMilestones(message)
}

func (m *Milestone) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	// we require roles for all commands in milestones
	return permutil.CheckModOrAbove(message.Roles)
}
