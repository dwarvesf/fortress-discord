package profile

import (
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
)

func (e *ProfileCmd) GetProfile(message *model.DiscordMessage) error {
	rawFormattedContent := stringutils.FormatString(message.RawContent)

	extractDiscordID := stringutils.ExtractPattern(rawFormattedContent, constant.DiscordIDRegexPattern)

	if len(extractDiscordID) == 0 || len(extractDiscordID) > 1 {
		return e.view.Error().Raise(message, "There is no valid user or more than one user tagged in your message.")
	}

	employee, err := e.svc.Profile().Get(extractDiscordID[0])
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get employee profile.")
	}

	return e.view.Profile().Get(message, employee)
}
