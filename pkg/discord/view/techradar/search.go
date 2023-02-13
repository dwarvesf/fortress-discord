package techradar

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *TechRadar) Search(original *model.DiscordMessage, topics []*model.TechRadarTopic) error {
	var content string

	for i := range topics {
		topic := topics[i]
		content += fmt.Sprintf("%d ・ %s - %s \n", i+1, topic.Name, topic.Ring)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepengoithien:940973057865027634> Topics <:pepengoithien:940973057865027634> ",
		Description: content,
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}
