package profile

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
	"strings"
)

func (e *ProfileCmd) Prefix() []string {
	return []string{"profile"}
}

// Execute is where we handle logic for each command
func (e *ProfileCmd) Execute(message *model.DiscordMessage) error {
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
	isChannelWhitelisted := false
	whiteListedChannels := strings.Split(e.cfg.Discord.WhiteListedChannels, ",")
	for _, id := range whiteListedChannels {
		if message.ChannelId == strings.TrimSpace(id) {
			isChannelWhitelisted = true
		}
	}

	if !isChannelWhitelisted {
		return false, []string{}
	}

	return permutil.CheckSmodOrAbove(message.Roles)
}
