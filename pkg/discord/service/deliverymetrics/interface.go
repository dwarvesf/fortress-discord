package deliverymetrics

import (
	"github.com/bwmarrin/discordgo"
)

type DeliveryMetricsServicer interface {
	GetWeeklyReportDiscordMsg() (*discordgo.MessageEmbed, error)
	GetMonthlyReportDiscordMsg() (*discordgo.MessageEmbed, error)

	SyncData() error
}
