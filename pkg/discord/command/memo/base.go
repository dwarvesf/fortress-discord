package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Memo) Prefix() []string {
	return []string{"memo", "memos"}
}

// Execute is where we handle logic for each command
func (a *Memo) Execute(message *model.DiscordMessage) error {
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

func (a *Memo) Name() string {
	return "Memo Command"
}

func (a *Memo) Help(message *model.DiscordMessage) error {
	return nil
}

func (a *Memo) DefaultCommand(message *model.DiscordMessage) error {
	return a.List(message)
}
