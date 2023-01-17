package techradar

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *TechRadar) ListAdopt(original *model.DiscordMessage, topics []*model.TechRadarTopic) error {
	var content string

	for i := range topics {
		topic := topics[i]
		content += fmt.Sprintf("%d ・ %s \n", i+1, topic.Name)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:okeoke:819507964414787584> Adopt Topics <:okeoke:819507964414787584> ",
		Description: content,
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}
