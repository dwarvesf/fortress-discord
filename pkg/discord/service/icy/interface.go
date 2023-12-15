package icy

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IcyServicer interface {
	GetWeeklyDistribution() ([]*model.Icy, error)
	ListUnpaidSalaryAdvances() (*model.SalaryAdvanceReport, error)
	GetIcyAccounting() (*model.IcyAccounting, error)
}
