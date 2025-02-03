package digest

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Digest) Prefix() []string {
	return []string{"digest"}
}

// Execute is where we handle logic for each command
func (a *Digest) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return a.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list", "ls":
		return a.List(message)
	case "help", "h":
		return a.Help(message)
	default:
		return a.Help(message)
	}
}

func (a *Digest) Name() string {
	return "Digest Command"
}

func (a *Digest) Help(message *model.DiscordMessage) error {
	return a.view.Digest().Help(message)
}

func (a *Digest) DefaultCommand(message *model.DiscordMessage) error {
	return a.Help(message)
}

func (a *Digest) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
