package techradar

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *TechRadar) LogTopicSuccess(original *model.DiscordMessage, topicName string) error {
	// send message
	msg := &discordgo.MessageEmbed{
		Title:       "Log topic success",
		Description: fmt.Sprintf("You have logged topic **%s** successfully", topicName),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func (e *TechRadar) LogTopicFailed(original *model.DiscordMessage, err string) error {
	// send message
	msg := &discordgo.MessageEmbed{
		Title:       "Log topic failed",
		Description: fmt.Sprintf("Failed with **%s** message", err),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
