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
	timePeriod := "30d"
	isLeaderboard := false

	if len(message.ContentArgs) > 1 {
		if message.ContentArgs[1] == "help" {
			return c.Help(message)
		}

		if message.ContentArgs[1] == "top" {
			timePeriod = "10y"
			isLeaderboard = true
			if len(message.ContentArgs) > 2 {
				timePeriod = strings.Join(message.ContentArgs[2:], "")
			}
		} else {
			userID := stringutils.ExtractDiscordID(message.ContentArgs[1])
			if userID != "" {
				if len(message.ContentArgs) > 2 {
					timePeriod = strings.Join(message.ContentArgs[2:], "")
				}
			} else {
				timePeriod = strings.Join(message.ContentArgs[1:], "")
			}
		}
	}

	after, timeAmount, timeUnit, err := stringutils.ParseTimePeriod(now, timePeriod)
	if err != nil {
		return err
	}

	if isLeaderboard {
		return c.GetOgifLeaderboard(message, *after, timeAmount, timeUnit)
	}

	return c.FetchOgifStats(message, message.Author.ID, *after, timeAmount, timeUnit)
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
