package techradar

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *TechRadar) ListAssess(original *model.DiscordMessage, topics []*model.TechRadarTopic) error {
	var content string

	for i := range topics {
		topic := topics[i]
		content += fmt.Sprintf("%d ・ %s \n", i+1, topic.Name)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<pepenote:885515949673951282> Assess Topics <pepenote:885515949673951282> ",
		Description: content,
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}
