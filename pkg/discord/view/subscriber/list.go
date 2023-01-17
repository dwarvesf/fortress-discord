package subscriber

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (s *Subscriber) ListNew(original *model.DiscordMessage, subs []*model.Subscriber) error {
	var content string
	content += fmt.Sprintf("Last 14 days has total %d new  subscribers \n\n", len(subs))
	for i := range subs {
		sub := subs[i]
		content += fmt.Sprintf("- %s\n", sub.Email)
	}

	msg := &discordgo.MessageEmbed{
		Title: "<:pepe_ping:1028964391690965012> Weekly Subscribers	<:pepe_ping:1028964391690965012>",
		Description: content,
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
