package subscriber

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (s *Subscriber) ListNew(original *model.DiscordMessage, subs []*model.Subscriber) error {
	var content string

	// group into 'source'
	sourceMap := make(map[string][]*model.Subscriber)
	for i := range subs {
		var source string
		if len(subs[i].Source) > 0 {
			source = subs[i].Source[0]
		} else {
			source = "Unknown Source"
		}
		sourceMap[source] = append(sourceMap[source], subs[i])
	}

	content += fmt.Sprintf("Last 14 days has total %d new subscribers \n\n", len(subs))
	for k, v := range sourceMap {
		content += fmt.Sprintf("**%s:**\n", k)
		for i := range v {
			content += fmt.Sprintf("- %s\n", v[i].Email)
		}
		content += "\n"
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepe_ping:1028964391690965012> Weekly Subscribers <:pepe_ping:1028964391690965012>",
		Description: content,
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
