package icy

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IcyViewer interface {
	List(original *model.DiscordMessage, earns []*model.Icy) error
	Accounting(original *model.DiscordMessage, icyAccounting *model.IcyAccounting, report *model.SalaryAdvanceReport) error
	PersonalInfo(original *model.DiscordMessage, accounting *model.IcyAccounting, totalEarned *model.ICYTotalEarned, earnedTxns []*model.ICYEarnedTransaction) error
	Help() error
}
