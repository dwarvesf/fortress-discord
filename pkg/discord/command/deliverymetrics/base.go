package deliverymetrics

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *DeliveryMetricsCmd) Prefix() []string {
	return []string{"delivery", "dlvy"}
}

// Execute is where we handle logic for each command
func (e *DeliveryMetricsCmd) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?sum`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	switch message.ContentArgs[1] {
	case "weekly":
		return e.WeeklyReport(message)
	case "monthly":
		return e.MonthlyReport(message)
	case "sync":
		return e.SyncRawData(message)
	case "help":
		return e.Help(message)
	default:
		return e.DefaultCommand(message)
	}
}

func (e *DeliveryMetricsCmd) Name() string {
	return "Delivery Metrics"
}

func (e *DeliveryMetricsCmd) Help(message *model.DiscordMessage) error {
	return e.view.DeliveryMetrics().Help(message)
}

func (e *DeliveryMetricsCmd) DefaultCommand(message *model.DiscordMessage) error {
	return e.Help(message)
}

func (e *DeliveryMetricsCmd) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	allowList := []string{
		"151497832853929986", // hanngo
		"567326528216760320", // hnh
		"790170208228212766", // thanh
		"361172853326086144", // huy tieu
		"797038457993297952", // nhut huynh
		"797042642600722473", // nam
		"794443008824049695", // tay
	}

	// check if user is in allow list
	for _, id := range allowList {
		if message.Author.ID == id {
			return true, []string{}
		}
	}

	return false, []string{}
}
