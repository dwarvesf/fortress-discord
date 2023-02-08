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

func (d *Done) Repost(original *model.DiscordMessage, msg string, channelId string, icy string) error {
	content := tagutil.FormatUser(original.Author.ID) + ": " + msg + "\n\n" + fmt.Sprintf("Thanks for your work, **%s** ICY has been sent to your balances", icy)
	embed := &discordgo.MessageEmbed{
		Description: content,
		Color:       5814783,
	}

	return base.SendEmbededMessageWithChannel(d.ses, original, embed, channelId)
}
