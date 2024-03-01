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

	//switch message.ContentArgs[1] {
	//case "with":
	//	return e.Home(message)
	//}

	return nil
}

func (e *Withdraw) Name() string {
	return "Home Command"
}

func (e *Withdraw) Help(message *model.DiscordMessage) error {
	return nil
}

func (e *Withdraw) DefaultCommand(message *model.DiscordMessage) error {
	return e.Withdraw(message)
}

func (e *Withdraw) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
