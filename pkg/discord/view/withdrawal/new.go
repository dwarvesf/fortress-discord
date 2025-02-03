package withdrawal

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Viewer interface {
	Home(original *model.DiscordMessage, in *model.WithdrawInput, banks []model.Bank) error
	ErrorWithdraw(original *model.DiscordMessage, err error) error
	//CompleteAdvanceSalary(original *model.DiscordMessage, salaryAdvance model.SalaryAdvance) error
	Help(original *model.DiscordMessage) error
}

type Withdraw struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &Withdraw{
		ses: ses,
	}
}
