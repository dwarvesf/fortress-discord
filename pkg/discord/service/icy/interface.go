package icy

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IcyServicer interface {
	GetWeeklyDistribution() ([]*model.Icy, error)
	ListUnpaidSalaryAdvances() (*model.SalaryAdvanceReport, error)
	GetIcyAccounting() (*model.IcyAccounting, error)
	ListICYEarnedTransactions(discordID string, page, size int) ([]*model.ICYEarnedTransaction, error)
	GetICYTotalEarned(discordID string) (*model.ICYTotalEarned, error)
	Get30daysTotalReward() (*model.ICYTotalEarned, error)
	GetBTCTreasury() (*model.IcyWeb3BigIntResponse, error)
	GetIcyRate() (*model.IcyWeb3BigIntResponse, error)
}
