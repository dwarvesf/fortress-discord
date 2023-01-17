package adopt

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Adopt) Prefix() []string {
	return []string{"adopt", "index"}
}

// Execute is where we handle logic for each command
func (a *Adopt) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return a.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":
		return a.List(message)
	}

	return nil
}

func (a *Adopt) Name() string {
	return "Hold Command"
}

func (a *Adopt) Help(message *model.DiscordMessage) error {
	return nil
}

func (a *Adopt) DefaultCommand(message *model.DiscordMessage) error {
	return a.List(message)
}
