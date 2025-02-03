package assess

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Assess) Prefix() []string {
	return []string{"assess"}
}

// Execute is where we handle logic for each command
func (a *Assess) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return a.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "list", "ls":
		return a.List(message)
	case "help", "h":
		return a.Help(message)
	default:
		return a.DefaultCommand(message)
	}
}

func (a *Assess) Name() string {
	return "Assess Command"
}

func (a *Assess) Help(message *model.DiscordMessage) error {
	return a.view.TechRadar().AssessHelp(message)
}

func (a *Assess) DefaultCommand(message *model.DiscordMessage) error {
	return a.Help(message)
}

func (a *Assess) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
