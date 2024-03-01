package deliverymetrics

import (
	"github.com/bwmarrin/discordgo"
)

type DeliveryMetricsServicer interface {
	GetWeeklyReportDiscordMsg() (*discordgo.MessageEmbed, error)
	GetMonthlyReportDiscordMsg(now bool) (*discordgo.MessageEmbed, error)

	SyncData() error
}
