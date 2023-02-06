package milestone

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Milestone) Prefix() []string {
	return []string{"milestone", "milestones"}
}

// Execute is where we handle logic for each command
func (a *Milestone) Execute(message *model.DiscordMessage) error {
	return a.ListMilestones(message)
}

func (a *Milestone) Name() string {
	return "Milestone Command"
}

func (a *Milestone) Help(message *model.DiscordMessage) error {
	return nil
}

func (a *Milestone) DefaultCommand(message *model.DiscordMessage) error {
	return a.ListMilestones(message)
}
