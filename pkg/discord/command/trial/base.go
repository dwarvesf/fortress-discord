package trial

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *Trial) Prefix() []string {
	return []string{"trial"}
}

// Execute is where we handle logic for each command
func (t *Trial) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return t.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "list", "ls":
		return t.List(message)
	case "help", "h":
		return t.Help(message)
	default:
		return t.Help(message)
	}
}

func (t *Trial) Name() string {
	return "Trial Command"
}

func (t *Trial) Help(message *model.DiscordMessage) error {
	return t.view.TechRadar().TrialHelp(message)
}

func (t *Trial) DefaultCommand(message *model.DiscordMessage) error {
	return t.Help(message)
}

func (t *Trial) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
