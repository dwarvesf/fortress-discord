package errors

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/tagutil"
)

func (e *Error) NotHavePermission(original *model.DiscordMessage, required []string) error {
	tmpl := "You don't have a required role to run this command \n This specific command require these following role: %s"
	var roles []string
	for i := range required {
		roles = append(roles, tagutil.FormatRole(required[i]))
	}
	msg := &discordgo.MessageEmbed{
		Title:       ":warning: **You can't execute this command** :warning: ",
		Description: fmt.Sprintf(tmpl, strings.Join(roles, " ,")),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
