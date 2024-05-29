package done

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (d *Done) Success(original *model.DiscordMessage) error {
	embed := &discordgo.MessageEmbed{
		Title:       "Success",
		Description: "Cmd has been executed successfully",
		Color:       0x00ff00,
	}

	return base.SendEmbededMessage(d.ses, original, embed)
}
