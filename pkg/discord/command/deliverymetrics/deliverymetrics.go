package deliverymetrics

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
)

func (e *DeliveryMetricsCmd) WeeklyReport(message *model.DiscordMessage) error {
	// 1. get data from service
	r, err := e.svc.DeliveryMetrics().GetWeeklyReportDiscordMsg()
	if err != nil {
		e.L.Error(err, "can't get WeeklyReport")
		return err
	}

	// 2. render
	return e.view.DeliveryMetrics().Send(message, r)
}

func (e *DeliveryMetricsCmd) MonthlyReport(message *model.DiscordMessage) error {
	rawFormattedContent := stringutils.FormatString(message.RawContent)

	// 1. get data from service
	times := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternTime)
	now := false
	if len(times) > 0 {
		now = times[0] == "now"
	}

	r, err := e.svc.DeliveryMetrics().GetMonthlyReportDiscordMsg(now)
	if err != nil {
		e.L.Error(err, "can't get MonthlyReport")
		return err
	}

	// 2. render
	return e.view.DeliveryMetrics().Send(message, r)
}

func (e *DeliveryMetricsCmd) SyncRawData(message *model.DiscordMessage) error {
	msg := &discordgo.MessageEmbed{
		Title:       "**Delivery Sync**",
		Description: "Sync data successfully \n",
	}

	err := e.svc.DeliveryMetrics().SyncData()
	if err != nil {
		e.L.Error(err, "can't get SyncRawData response")
		msg.Description = "Sync data failed " + err.Error() + "\n"
	}

	// 2. render
	return e.view.DeliveryMetrics().Send(message, msg)
}
