package deliverymetrics

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c *DeliveryMetricsCmd) WeeklyReport(message *model.DiscordMessage) error {
	// 1. get data from service
	r, err := c.svc.DeliveryMetrics().GetWeeklyReportDiscordMsg()
	if err != nil {
		c.L.Error(err, "can't get list of Memo")
		return err
	}

	// 2. render
	return c.view.DeliveryMetrics().Send(message, r)
}

func (c *DeliveryMetricsCmd) MonthlyReport(message *model.DiscordMessage) error {
	// 1. get data from service
	r, err := c.svc.DeliveryMetrics().GetMonthlyReportDiscordMsg()
	if err != nil {
		c.L.Error(err, "can't get list of Memo")
		return err
	}

	// 2. render
	return c.view.DeliveryMetrics().Send(message, r)
}
