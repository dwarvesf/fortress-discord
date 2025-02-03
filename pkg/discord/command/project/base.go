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
		"1136175977236549693", "1064460652720160808", // dev-channel
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

	// handle command for 2 args input from user, e.g `?project list`
	switch message.ContentArgs[1] {
	case "pnl":
		return e.GetProjectPnL(message)
	case "commission", "com":
		return e.GetProjectCommissionModels(message)
	case "list", "ls":
		return e.GetProjectList(message)
	case "help", "h":
		return e.Help(message)
	default:
		return e.Help(message)
	}
}

func (e *ProjectCmd) Name() string {
	return "Project"
}

func (e *ProjectCmd) Help(message *model.DiscordMessage) error {
	return e.view.Project().Help(message)
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

	return permutil.CheckModOrAbove(roles)
}
