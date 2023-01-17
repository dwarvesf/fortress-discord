package new

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (n *NewCommand) Prefix() []string {
	return []string{"new"}
}

// Execute is where we handle logic for each command
func (n *NewCommand) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return n.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return n.List(message)
	}

	return nil
}

func (n *NewCommand) Name() string {
	return "New Command"
}

func (n *NewCommand) Help(message *model.DiscordMessage) error {
	return nil
}

func (n *NewCommand) DefaultCommand(message *model.DiscordMessage) error {
	return n.List(message)
}
