package withdraw

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Withdraw struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) Commander {
	return &Withdraw{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Withdraw) Withdraw(message *model.DiscordMessage) error {
	cond, err := e.svc.Withdrawal().CheckWithdrawCondition(message.Author.ID)
	if err != nil {
		e.L.Error(err, "failed to check withdraw condition")
		return e.view.Withdraw().ErrorWithdraw(message, err)
	}

	banks, err := e.svc.Withdrawal().GetBanks("", "", "")
	if err != nil {
		e.L.Error(err, "failed to get banks")
		return e.view.Withdraw().ErrorWithdraw(message, err)
	}

	in := &model.WithdrawInput{
		ICYAmount:  cond.ICYAmount,
		VNDAmount:  cond.VNDAmount,
		ICYVNDRate: cond.ICYVNDRate,
	}
	return e.view.Withdraw().Home(message, in, banks)
}
