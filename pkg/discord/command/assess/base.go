package assess

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *Assess) Prefix() []string {
	return []string{"assess"}
}

// Execute is where we handle logic for each command
func (t *Assess) Execute(message *model.DiscordMessage) error {
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

func (t *Assess) Name() string {
	return "Assess Command"
}

func (t *Assess) Help(message *model.DiscordMessage) error {
	return nil
}

func (t *Assess) DefaultCommand(message *model.DiscordMessage) error {
	return t.List(message)
}

func (a *Assess) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
