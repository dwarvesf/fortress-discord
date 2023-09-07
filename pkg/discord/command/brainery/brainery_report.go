package brainery

import (
	"strings"
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

const dateRegexPattern = `(\d{4}-\d{2}-\d{2})`

func (e *Brainery) Report(message *model.DiscordMessage) error {
	rawFormattedContent := formatString(message.RawContent)
	args := strings.Split(rawFormattedContent, " ")
	targetChannelID := message.ChannelId
	defaultQueryDate := time.Now().Format("2006-01-02")

	reportView := "weekly"
	if len(args) > 2 {
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
	extractDate := extractPattern(rawFormattedContent, dateRegexPattern)
	if len(extractDate) > 0 {
		parsedDate, err := time.Parse("2006-01-02", extractDate[0])
		if err != nil {
			return e.view.Error().Raise(message, "The date format is invalid.")
		}

		defaultQueryDate = parsedDate.Format("2006-01-02")
	}

	braineryData, err := e.svc.Brainery().Report(reportView, defaultQueryDate)
	if err != nil {
		return e.view.Error().Raise(message, err.Error())
	}

	return e.view.Brainery().Report(message, reportView, braineryData, targetChannelID)
}
