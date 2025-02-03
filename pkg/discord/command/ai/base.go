package ai

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *AI) Prefix() []string {
	return []string{"ai"}
}

// Execute is where we handle logic for each command
func (a *AI) Execute(message *model.DiscordMessage) error {
	if len(message.ContentArgs) == 2 {
		switch message.ContentArgs[1] {
		case "help", "h":
			return a.Help(message)
		}
	}

	return a.DefaultCommand(message)
}

func (a *AI) Name() string {
	return "AI Command"
}

func (a *AI) Help(message *model.DiscordMessage) error {
	return a.view.AI().Help(message)
}

func (a *AI) DefaultCommand(message *model.DiscordMessage) error {
	return a.ProcessAI(message)
}

func (a *AI) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
