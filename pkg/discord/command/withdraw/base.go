package withdraw

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Withdraw) Prefix() []string {
	return []string{"withdraw"}
}

// Execute is where we handle logic for each command
func (e *Withdraw) Execute(message *model.DiscordMessage) error {
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "help", "h":
		return e.Help(message)
	}

	return e.DefaultCommand(message)
}

func (e *Withdraw) Name() string {
	return "Withdraw Command"
}

func (e *Withdraw) Help(message *model.DiscordMessage) error {
	return e.view.Withdraw().Help(message)
}

func (e *Withdraw) DefaultCommand(message *model.DiscordMessage) error {
	return e.Withdraw(message)
}

func (e *Withdraw) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
