package profile

import (
	"fmt"
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/profile"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
	"strings"
)

func (e *ProfileCmd) GetProfile(message *model.DiscordMessage) error {
	rawFormattedContent := stringutils.FormatString(message.RawContent)
	fmt.Println(rawFormattedContent)
	extractDiscordID := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternDiscordID)
	extractEmail := stringutils.ExtractEmailPattern(rawFormattedContent)

	in := profile.EmployeeSearch{}
	if len(extractDiscordID) == 1 {
		in.DiscordID = extractDiscordID[0]
	} else {
		discordID := stringutils.ExtractNumber(rawFormattedContent)
		if len(discordID) > 0 {
			in.DiscordID = discordID[0]
		}
	}

	if len(extractEmail) == 1 {
		in.Email = extractEmail[0]
	}

	if in.DiscordID == "" && in.Email == "" {
		in.Key = strings.Split(rawFormattedContent, " ")[1]
	}

	employees, err := e.svc.Profile().GetEmployeeList(in)
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get employee profile.")
	}
	fmt.Println(employees)

	return e.view.Profile().List(message, employees)
}
