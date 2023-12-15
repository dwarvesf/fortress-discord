package icy

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IcyViewer interface {
	List(original *model.DiscordMessage, earns []*model.Icy) error
	Accounting(original *model.DiscordMessage, icyAccounting *model.IcyAccounting, report *model.SalaryAdvanceReport) error
	Help() error
}
