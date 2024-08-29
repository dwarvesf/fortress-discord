package ogif

import (
	"strings"
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
)

func (c command) Prefix() []string {
	return []string{"ogif"}
}

// Execute is where we handle logic for each command
func (c command) Execute(message *model.DiscordMessage) error {
	now := time.Now()
	userID := message.Author.ID
	timePeriod := "30d"
	isUserOmitted := false

	if len(message.ContentArgs) > 1 {
		if message.ContentArgs[1] == "help" {
			return c.Help(message)
		}

		extractedID := stringutils.ExtractDiscordID(message.ContentArgs[1])
		if extractedID != "" {
			userID = extractedID
		} else {
			isUserOmitted = true
			timePeriod = strings.Join(message.ContentArgs[1:], "")
		}
	}

	if len(message.ContentArgs) > 2 && !isUserOmitted {
		timePeriod = strings.Join(message.ContentArgs[2:], "")
	}

	after, timeAmount, timeUnit, err := stringutils.ParseTimePeriod(now, timePeriod)
	if err != nil {
		return err
	}

	return c.FetchOgifStats(message, userID, *after, timeAmount, timeUnit)
}

func (c command) Name() string {
	return "OGIF command"
}

func (c command) Help(message *model.DiscordMessage) error {
	return c.view.Ogif().Help(message)
}

func (c command) DefaultCommand(message *model.DiscordMessage) error {
	return nil
}

func (c command) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}

func (c command) ChannelPermissionCheck(message *model.DiscordMessage) bool {
	return true
}
