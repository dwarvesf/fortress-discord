package staff

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *Staff) Prefix() []string {
	return []string{"staff"}
}

// Execute is where we handle logic for each command
func (t *Staff) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return t.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return t.List(message)
	}

	return nil
}

func (t *Staff) Name() string {
	return "Staff Command"
}

func (t *Staff) Help(message *model.DiscordMessage) error {
	return nil
}

func (t *Staff) DefaultCommand(message *model.DiscordMessage) error {
	return t.List(message)
}
