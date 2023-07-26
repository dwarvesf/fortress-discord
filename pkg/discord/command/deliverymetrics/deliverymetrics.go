package deliverymetrics

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c *DeliveryMetricsCmd) WeeklyReport(message *model.DiscordMessage) error {
	// 1. get data from service
	r, err := c.svc.DeliveryMetrics().GetWeeklyReportDiscordMsg()
	if err != nil {
		c.L.Error(err, "can't get WeeklyReport")
		return err
	}

	// 2. render
	return c.view.DeliveryMetrics().Send(message, r)
}

func (c *DeliveryMetricsCmd) MonthlyReport(message *model.DiscordMessage) error {
	// 1. get data from service
	r, err := c.svc.DeliveryMetrics().GetMonthlyReportDiscordMsg()
	if err != nil {
		c.L.Error(err, "can't get MonthlyReport")
		return err
	}

	// 2. render
	return c.view.DeliveryMetrics().Send(message, r)
}

func (c *DeliveryMetricsCmd) SyncRawData(message *model.DiscordMessage) error {
	msg := &discordgo.MessageEmbed{
		Title:       "**Delivery Sync**",
		Description: "Sync data successfully \n",
	}

	err := c.svc.DeliveryMetrics().SyncData()
	if err != nil {
		c.L.Error(err, "can't get SyncRawData response")
		msg.Description = "Sync data failed " + err.Error() + "\n"
	}

	// 2. render
	return c.view.DeliveryMetrics().Send(message, msg)
}
