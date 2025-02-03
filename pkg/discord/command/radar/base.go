package radar

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (h *Radar) Prefix() []string {
	return []string{"radar"}
}

// Execute is where we handle logic for each command
func (h *Radar) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return h.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list", "ls":
		return h.List(message)
	case "log":
		return h.Log(message)
	case "help", "h":
		return h.Help(message)
	default:
		return h.DefaultCommand(message)
	}
}

func (h *Radar) Name() string {
	return "Radar Command"
}

func (h *Radar) Help(message *model.DiscordMessage) error {
	return h.view.TechRadar().AdoptHelp(message)
}

func (h *Radar) DefaultCommand(message *model.DiscordMessage) error {
	return h.Help(message)
}

func (h *Radar) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
