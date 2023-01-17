package earn

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Earn) Prefix() []string {
	return []string{"earn"}
}

// Execute is where we handle logic for each command
func (e *Earn) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return e.List(message)
	}

	return nil
}

func (e *Earn) Name() string {
	return "Earn Command"
}

func (e *Earn) Help() {
}

func (e *Earn) DefaultCommand(message *model.DiscordMessage) error {
	return e.List(message)
}
