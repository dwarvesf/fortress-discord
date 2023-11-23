package salary

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Salary struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) SalaryCommander {
	return &Salary{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Salary) Advance(message *model.DiscordMessage) error {
	salaryAdvance, err := e.svc.Salary().CheckAdvanceSalary(message.Author.ID)
	if err != nil {
		return err
	}
	return e.view.Salary().EnterAmountAdvanceSalary(message, *salaryAdvance)
}
