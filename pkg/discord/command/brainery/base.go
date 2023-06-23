package brainery

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Brainery) Prefix() []string {
	return []string{"brainery"}
}

// Execute is where we handle logic for each command
func (e *Brainery) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?sum`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "post":
		return e.Post(message)
	case "report":
		return e.Report(message)
	default:
		return e.DefaultCommand(message)
	}
}

func (e *Brainery) Name() string {
	return "Brainery Command"
}

func (e *Brainery) Help(message *model.DiscordMessage) error {
	return e.view.Brainery().Help(message)
}

func (e *Brainery) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *Brainery) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	allowList := []string{
		"151497832853929986", //hanngo
		"567326528216760320", //hnh
		"790170208228212766", //thanh
		"184354519726030850", //tom
		"797042642600722473", //nam
	}

	// check if user is in allow list
	for _, id := range allowList {
		if message.Author.ID == id {
			return true, []string{}
		}
	}

	return false, []string{}
}
