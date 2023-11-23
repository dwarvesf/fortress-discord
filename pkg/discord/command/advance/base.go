package salary

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Salary) Prefix() []string {
	return []string{"salary"}
}

// Execute is where we handle logic for each command
func (e *Salary) Execute(message *model.DiscordMessage) error {
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "advance":
		return e.Advance(message)
	}

	return nil
}

func (e *Salary) Name() string {
	return "Salary Command"
}

func (e *Salary) Help(message *model.DiscordMessage) error {
	return nil
}

func (e *Salary) DefaultCommand(message *model.DiscordMessage) error {
	return e.Advance(message)
}

func (e *Salary) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
