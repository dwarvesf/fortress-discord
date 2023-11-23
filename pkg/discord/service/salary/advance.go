package salary

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Salary struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) SalaryServicer {
	return &Salary{
		adapter: adapter,
		l:       l,
	}
}

// SalaryAdvance implements SalaryServicer.
func (s *Salary) CheckAdvanceSalary(discordID string) (*model.CheckSalaryAdvance, error) {
	salaryAdvance, err := s.adapter.Fortress().CheckAdvanceSalary(discordID)
	if err != nil {
		s.l.Error(err, "can't check advance salary")
		return nil, err
	}
	return &salaryAdvance.Data, nil
}

// SalaryAdvance implements SalaryServicer.
func (s *Salary) SalaryAdvance(discordID, amount string) (*model.SalaryAdvance, error) {
	salaryAdvance, err := s.adapter.Fortress().SalaryAdvance(discordID, amount)
	if err != nil {
		s.l.Error(err, "can't check advance salary")
		return nil, err
	}
	return &salaryAdvance.Data, nil
}
