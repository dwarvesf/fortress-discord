package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c command) Prefix() []string {
	return []string{"news"}
}

// Execute is where we handle logic for each command
func (c command) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, c.g `?earn`
	if len(message.ContentArgs) == 1 {
		return c.DefaultCommand(message)
	}

	// handle command for 2 args input from user, c.g `?earn list`
	switch message.ContentArgs[1] {
	case "help":
		return c.Help(message)
	case "reddit":
		switch len(message.ContentArgs) {
		case 2:
			return c.DefaultCommand(message)
		default:
			return c.Reddit(message, message.ContentArgs[2])
		}
	}

	return nil
}

func (c command) Name() string {
	return "News Command"
}

func (c command) Help(message *model.DiscordMessage) error {
	return c.view.News().Help(message)
}

func (c command) DefaultCommand(message *model.DiscordMessage) error {
	return c.Help(message)
}

func (c command) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}

func (c command) ChannelPermissionCheck(message *model.DiscordMessage) bool {
	return true
}
