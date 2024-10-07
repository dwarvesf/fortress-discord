package profile

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (e *ProfileCmd) Prefix() []string {
	return []string{"profile"}
}

// Execute is where we handle logic for each command
func (e *ProfileCmd) Execute(message *model.DiscordMessage) error {
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

	// default command for only 1 args input from user, e.g `?profile`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	return e.GetProfile(message)
}

func (e *ProfileCmd) Name() string {
	return "Profile"
}

func (e *ProfileCmd) Help(message *model.DiscordMessage) error {
	return e.view.DeliveryMetrics().Help(message)
}

func (e *ProfileCmd) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *ProfileCmd) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	if e.cfg.Env != "prod" {
		return true, []string{}
	}

	return permutil.CheckSmodOrAbove(message.Roles)
}
