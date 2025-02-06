package index

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (i *Index) Prefix() []string {
	return []string{"index"}
}

// Execute is where we handle logic for each command
func (i *Index) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return i.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "list", "ls":
		return i.Search(message)
	case "help", "h":
		return i.Help(message)
	default:
		return i.DefaultCommand(message)
	}
}

func (i *Index) Name() string {
	return "Index Command"
}

func (i *Index) Help(message *model.DiscordMessage) error {
	return i.view.TechRadar().IndexHelp(message)
}

func (i *Index) DefaultCommand(message *model.DiscordMessage) error {
	return i.Help(message)
}

func (i *Index) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
