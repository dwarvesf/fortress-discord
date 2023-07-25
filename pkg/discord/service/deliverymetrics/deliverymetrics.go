package deliverymetrics

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type DeliveryMetricSvc struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) DeliveryMetricsServicer {
	return &DeliveryMetricSvc{
		adapter: adapter,
		l:       l,
	}
}

func (e *DeliveryMetricSvc) GetWeeklyReportDiscordMsg() (*discordgo.MessageEmbed, error) {
	// get response from fortress
	reportAdapter, err := e.adapter.Fortress().GetDeliveryMetricsWeeklyReportDiscordMsg()
	if err != nil {
		e.l.Error(err, "can't get open memo from fortress")
		return nil, err
	}

	return reportAdapter.Data, nil
}

func (e *DeliveryMetricSvc) GetMonthlyReportDiscordMsg() (*discordgo.MessageEmbed, error) {
	// get response from fortress
	reportAdapter, err := e.adapter.Fortress().GetDeliveryMetricsMonthlyReportDiscordMsg()
	if err != nil {
		e.l.Error(err, "can't get open memo from fortress")
		return nil, err
	}

	return reportAdapter.Data, nil
}
