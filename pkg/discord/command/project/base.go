package project

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (e *ProjectCmd) Prefix() []string {
	return []string{"project"}
}

// Execute is where we handle logic for each command
func (e *ProjectCmd) Execute(message *model.DiscordMessage) error {
	allowedList := []string{
		"1072722777687199744",
		"1177538072531980318",
		"1136175977236549693", // dev-channel
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

	return e.GetProjectCommissionModels(message)
}

func (e *ProjectCmd) Name() string {
	return "Project"
}

func (e *ProjectCmd) Help(message *model.DiscordMessage) error {
	return e.view.DeliveryMetrics().Help(message)
}

func (e *ProjectCmd) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *ProjectCmd) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	if e.cfg.Env != "prod" {
		return true, []string{}
	}

	userRoles, err := e.svc.Profile().GetDiscordRoles(e.cfg.Discord.ID.DwarvesGuild, message.Author.ID)
	if err != nil {
		return false, []string{}
	}

	var roles []string
	roles = append(roles, userRoles...)
	roles = append(roles, message.Roles...)

	return permutil.CheckSmodOrAbove(roles)
}
