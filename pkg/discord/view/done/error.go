package done

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/tagutil"
)

func (d *Done) MissingContent(original *model.DiscordMessage) error {
	embed := &discordgo.MessageEmbed{
		Title:       "Missing content",
		Description: "Please make sure you input the correct message \n\n e.g: \n - ?done helping a friend with their API \n - ?done init project repositories",
	}

	return base.SendEmbededMessage(d.ses, original, embed)
}

func (d *Done) CantSendReward(original *model.DiscordMessage) error {
	embed := &discordgo.MessageEmbed{
		Title:       "Can't send reward",
		Description: fmt.Sprintf("I can't execute the transfer atm, please notify %s", tagutil.FormatUser(tagutil.HnhDiscordId)),
	}

	return base.SendEmbededMessage(d.ses, original, embed)
}
