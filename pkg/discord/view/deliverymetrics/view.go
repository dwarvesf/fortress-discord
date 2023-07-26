package deliverymetrics

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DeliveryMetrics struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) DeliveryMetricsViewer {
	return &DeliveryMetrics{
		ses: ses,
	}
}

func (v *DeliveryMetrics) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?delivery <type>** ・get delivery metrics report by week/month.",
		"**?delivery sync To sync data",
		"**?dlvy <type>** ・get delivery metrics report by week/month.",
		"*Example:* `?delivery weekly`",
		"*Example:* `?dlvy monthly`",
		"*Example:* `?dlvy sync`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}

func (v *DeliveryMetrics) Send(original *model.DiscordMessage, msg *discordgo.MessageEmbed) error {
	return base.SendEmbededMessage(v.ses, original, msg)
}
