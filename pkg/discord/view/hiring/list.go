package hiring

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Hiring) OpeningList(original *model.DiscordMessage, positions []*model.HiringPosition) error {
	var content string
	for i := range positions {
		content += fmt.Sprintf("%s ・ [Apply Now](%s)\n", positions[i].Name, "https://dwarves.notion.site/512210719fec4152bd76c87d3cae2d52?v=64e31aa8ed4b4ca3be2454a7e5db7b1a&p="+strings.Replace(positions[i].Id, "-", "", -1))
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeyes:885513213431648266> Open Positions <:pepeyes:885513213431648266> \n",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
