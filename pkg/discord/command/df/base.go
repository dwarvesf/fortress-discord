package df

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/permutil"
)

func (a *DF) Prefix() []string {
	return []string{"df"}
}

// Execute is where we handle logic for each command
func (d *DF) Execute(message *model.DiscordMessage) error {
	if len(message.ContentArgs) == 1 {
		return d.Help(message)
	}

	switch message.ContentArgs[1] {
	case "help", "h":
		return d.Help(message)
	}

	return d.DefaultCommand(message)
}

func (d *DF) Name() string {
	return "DF Command"
}

func (d *DF) Help(message *model.DiscordMessage) error {
	return d.view.DF().Help(message)
}

func (d *DF) DefaultCommand(message *model.DiscordMessage) error {
	return d.ProcessWithN8N(message)
}

func (d *DF) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	if d.cfg.Env != "prod" {
		return true, []string{}
	}

	userRoles, err := d.svc.Profile().GetDiscordRoles(d.cfg.Discord.ID.DwarvesGuild, message.Author.ID)
	if err != nil {
		return false, []string{}
	}

	var roles []string
	roles = append(roles, userRoles...)
	roles = append(roles, message.Roles...)

	return permutil.CheckAdmin(roles)
}
