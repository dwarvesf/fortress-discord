package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c command) Prefix() []string {
	return []string{"news"}
}

// Execute is where we handle logic for each command
func (c command) Execute(message *model.DiscordMessage) error {
	switch len(message.ContentArgs) {
	case 1:
		return c.DefaultCommand(message)
	case 2:
		switch message.ContentArgs[1] {
		case "help":
			return c.Help(message)
		default:
			return c.DefaultCommand(message)
		}
	case 3:
		return c.Fetch(message, message.ContentArgs[1], message.ContentArgs[2])
	default:
		return c.DefaultCommand(message)
	}
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
