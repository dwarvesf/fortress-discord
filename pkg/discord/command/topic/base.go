package topic

import "github.com/dwarvesf/fortress-discord/pkg/model"

func (e *Topic) Prefix() []string {
	return []string{"topic", "topics"}
}

// Execute is where we handle logic for each command
func (e *Topic) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn sum`
	switch message.ContentArgs[1] {
	case "list", "l":
		return e.List(message)
	case "help", "h":
		return e.Help(message)
	default:
		return e.DefaultCommand(message)
	}
}

func (e *Topic) Name() string {
	return "Topic Command"
}

func (e *Topic) Help(message *model.DiscordMessage) error {
	return e.view.Topic().Help(message)
}

// DefaultCommand handles the default command
func (e *Topic) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *Topic) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
