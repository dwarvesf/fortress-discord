package ai

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *AI) Prefix() []string {
	return []string{"ai"}
}

// Execute is where we handle logic for each command
func (a *AI) Execute(message *model.DiscordMessage) error {
	// For AI command, we always process the input, regardless of the number of arguments
	return a.DefaultCommand(message)
}

func (a *AI) Name() string {
	return "AI Command"
}

func (a *AI) Help(message *model.DiscordMessage) error {
	helpText := "The AI command processes any text input using an AI model.\n\n" +
		"Usage: ?ai <your text here>\n" +
		"Example: ?ai What is the capital of France?"

	return a.view.Error().Raise(message, helpText)
}

func (a *AI) DefaultCommand(message *model.DiscordMessage) error {
	return a.ProcessAI(message)
}

func (a *AI) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
