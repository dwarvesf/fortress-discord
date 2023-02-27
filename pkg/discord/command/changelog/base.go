package changelog

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (a *Changelog) Prefix() []string {
	return []string{"changelog"}
}

// Execute is where we handle logic for each command
func (a *Changelog) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return a.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "send":
		return a.Send(message)
	}

	return nil
}

func (a *Changelog) Name() string {
	return "Changelog Command"
}

func (a *Changelog) Help(message *model.DiscordMessage) error {
	return nil
}

func (a *Changelog) DefaultCommand(message *model.DiscordMessage) error {
	return nil
}

func (a *Changelog) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	allowList := []string{
		"151497832853929986", //hanngo
		"567326528216760320", //hnh
		"361172853326086144", //huytq
		"796991130184187944", //nikki
		"538931010688122881", //minhlq
	}

	// check if user is in allow list
	for _, id := range allowList {
		if message.Author.ID == id {
			return true, []string{}
		}
	}

	return false, []string{}
}
