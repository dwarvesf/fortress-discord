package deliverymetrics

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DeliveryMetricsViewer interface {
	Help(message *model.DiscordMessage) error

	Send(original *model.DiscordMessage, msg *discordgo.MessageEmbed) error
}
