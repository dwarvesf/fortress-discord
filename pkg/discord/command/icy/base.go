package icy

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Icy) Prefix() []string {
	return []string{"icy"}
}

// Execute is where we handle logic for each command
func (e *Icy) Execute(message *model.DiscordMessage) error {
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

func (e *Icy) Name() string {
	return "Icy Command"
}

func (e *Icy) Help(message *model.DiscordMessage) error {
	return nil
}

func (e *Icy) DefaultCommand(message *model.DiscordMessage) error {
	return e.List(message)
}

func (e *Icy) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
