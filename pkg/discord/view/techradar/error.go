package techradar

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *TechRadar) SearchEmpty(original *model.DiscordMessage) error {
	var content string = "No topics found, consider trying with another name or submit a new topic"

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepengoithien:940973057865027634> No Topics Found <:pepengoithien:940973057865027634> ",
		Description: content,
	}

	return base.SendEmbededMessage(t.ses, original, msg)
}
