package hold

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (h *Hold) Prefix() []string {
	return []string{"hold"}
}

// Execute is where we handle logic for each command
func (h *Hold) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return h.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return h.List(message)
	}

	return nil
}

func (h *Hold) Name() string {
	return "Hold Command"
}

func (h *Hold) Help(message *model.DiscordMessage) error {
	return nil
}

func (h *Hold) DefaultCommand(message *model.DiscordMessage) error {
	return h.List(message)
}
