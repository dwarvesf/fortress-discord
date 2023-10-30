package trend

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (t *Trend) Prefix() []string {
	return []string{"trend"}
}

// Execute is where we handle logic for each command
func (t *Trend) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?trend`
	if len(message.ContentArgs) == 1 {
		return t.Help(message)
	}

	// handle command for 2 args input from user, e.g `?earn sum`
	// switch message.ContentArgs[1] {
	// case "sum":
	// 	return e.Sum(message)
	// }

	return t.Trend(message)
}

func (t *Trend) Name() string {
	return "Trend Command"
}

func (t *Trend) Help(message *model.DiscordMessage) error {
	return t.view.Trend().Help(message)
}

// Default: golang, english, daily
func (t *Trend) DefaultCommand(message *model.DiscordMessage) error {
	return t.DefaultTrend(message)
}

func (e *Trend) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
