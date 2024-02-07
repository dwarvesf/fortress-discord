package icy

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Icy struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) IcyCommander {
	return &Icy{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

func (e *Icy) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Icy().GetWeeklyDistribution()
	if err != nil {
		e.L.Error(err, "can't get list of weekly icy distribution")
		return err
	}

	// 2. render
	return e.view.Icy().List(message, data)
}

func (e *Icy) Accounting(message *model.DiscordMessage) error {
	// 1. get data from service
	// 1.1 Get icy accounting info
	icyAccounting, err := e.svc.Icy().GetIcyAccounting()
	if err != nil {
		e.L.Error(err, "can't get icy accounting info")
		return err
	}

	// 1.2 Get list of unpaid salary advances
	report, err := e.svc.Icy().ListUnpaidSalaryAdvances()
	if err != nil {
		e.L.Error(err, "can't get list of unpaid salary advances")
		return err
	}

	// 1.3 Get total reward 30 days
	total30DaysReward, err := e.svc.Icy().Get30daysTotalReward()
	if err != nil {
		e.L.Error(err, "can't get total reward 30 days")
		return err
	}

	// 2. render
	return e.view.Icy().Accounting(message, icyAccounting, report, total30DaysReward)
}

func (e *Icy) PersonalInfo(message *model.DiscordMessage) error {
	// 1. get data from service
	// 1.1 Get icy accounting info
	icyAccounting, err := e.svc.Icy().GetIcyAccounting()
	if err != nil {
		e.L.Error(err, "can't get icy accounting info")
		return err
	}

	// 1.2 Get total icy earned
	totalEarned, err := e.svc.Icy().GetICYTotalEarned(message.Author.ID)
	if err != nil {
		e.L.Error(err, "can't get total icy earned")
		return err
	}

	// 1.3 Get list last 5 icy transactions
	earnedTxns, err := e.svc.Icy().ListICYEarnedTransactions(message.Author.ID, 0, 5)
	if err != nil {
		e.L.Error(err, "can't get list of last 5 icy transactions")
		return err
	}

	// 2. render
	return e.view.Icy().PersonalInfo(message, icyAccounting, totalEarned, earnedTxns)
}
