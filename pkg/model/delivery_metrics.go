package model

import "github.com/bwmarrin/discordgo"

type AdapterDeliveryMetricsReportMsg struct {
	Data *discordgo.MessageEmbed `json:"data"`
}
