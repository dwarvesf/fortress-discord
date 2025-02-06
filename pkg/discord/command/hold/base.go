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

	switch message.ContentArgs[1] {
	case "list", "ls":
		return h.List(message)
	case "help", "h":
		return h.Help(message)
	default:
		return h.DefaultCommand(message)
	}
}

func (h *Hold) Name() string {
	return "Hold Command"
}

func (h *Hold) Help(message *model.DiscordMessage) error {
	return h.view.TechRadar().HoldHelp(message)
}

func (h *Hold) DefaultCommand(message *model.DiscordMessage) error {
	return h.Help(message)
}

func (h *Hold) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
