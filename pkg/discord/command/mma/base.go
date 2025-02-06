package mma

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (e *MMACmd) Prefix() []string {
	return []string{"mma"}
}

// Execute is where we handle logic for each command
func (e *MMACmd) Execute(message *model.DiscordMessage) error {
	allowedList := []string{
		"1072722777687199744",
	}

	whiteListedChannels := strings.Split(e.cfg.Discord.WhiteListedChannels, ",")
	allowedList = append(allowedList, whiteListedChannels...)
	isChannelWhitelisted := false
	for _, id := range allowedList {
		if message.ChannelId == strings.TrimSpace(id) {
			isChannelWhitelisted = true
		}
	}

	if !isChannelWhitelisted {
		return e.view.Error().Raise(message, "This command is not allowed in this channel.")

	}

	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "template":
		return e.ExportTemplate(message)
	case "help", "h":
		return e.Help(message)
	default:
		return e.DefaultCommand(message)
	}
}

func (e *MMACmd) Name() string {
	return "MMA"
}

func (e *MMACmd) Help(message *model.DiscordMessage) error {
	return e.view.MMA().Help(message)
}

func (e *MMACmd) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *MMACmd) PermissionCheck(message *model.DiscordMessage) (bool, []string) {

	if e.cfg.Env != "prod" {
		return true, nil
	}

	return permutil.CheckSmodOrAbove(message.Roles)
}
