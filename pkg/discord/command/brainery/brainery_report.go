package brainery

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Brainery) Report(message *model.DiscordMessage) error {
	rawFormattedContent := formatString(message.RawContent)
	args := strings.Split(rawFormattedContent, " ")
	targetChannelID := message.ChannelId

	reportView := "weekly"
	if len(args) > 1 {
		reportView = args[2]
	}

	if !(reportView == "weekly" || reportView == "monthly") {
		return e.view.Error().Raise(message, "Report view should be weekly or monthly")
	}

	extractChannelID := extractPattern(rawFormattedContent, discordChannelIDRegexPattern)
	if len(extractChannelID) > 1 {
		return e.view.Error().Raise(message, "There is more than one target channel in your message.")
	}

	if len(extractChannelID) == 1 {
		targetChannelID = extractChannelID[0]
	}

	result, err := e.svc.Brainery().Report(reportView)
	if err != nil {
		return e.view.Error().Raise(message, err.Error())
	}

	return e.view.Brainery().Report(message, reportView, result, targetChannelID)
}
