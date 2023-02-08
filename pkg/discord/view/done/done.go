package done

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/tagutil"
)

type Done struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) DoneViewer {
	return &Done{
		ses: ses,
	}
}

func (d *Done) Repost(original *model.DiscordMessage, msg string, channelId string) error {
	embed := &discordgo.MessageEmbed{
		Description: tagutil.FormatUser(original.Author.ID) + ": " + msg,
	}

	return base.SendEmbededMessageWithChannel(d.ses, original, embed, channelId)
}

func (d *Done) NotifyIcyReward(original *model.DiscordMessage, msg string, icy string) error {
	tmpl :=
		`Thanks for your work, **%s** ICY has been sent to your balances

		Try $balances to see your rewards.
	`

	embed := &discordgo.MessageEmbed{
		Title:       ":ice_cube: You just got from ice :ice_cube:",
		Description: fmt.Sprintf(tmpl, icy),
	}

	return base.SendEmbededMessage(d.ses, original, embed)
}
