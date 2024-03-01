package withdrawal

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Withdrawal struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) Servicer {
	return &Withdrawal{
		adapter: adapter,
		l:       l,
	}
}

// CheckWithdrawCondition means check condition to withdraw money from ICY
func (s *Withdrawal) CheckWithdrawCondition(discordID string) (*model.CheckWithdrawCondition, error) {
	rs, err := s.adapter.Fortress().CheckWithdrawCondition(discordID)
	if err != nil {
		s.l.Error(err, "failed to check withdraw condition")
		return nil, err
	}
	return &rs.Data, nil
}

func (s *Withdrawal) GetBanks(id, bin, swiftCode string) ([]model.Bank, error) {
	rs, err := s.adapter.Fortress().GetBanks(id, bin, swiftCode)
	if err != nil {
		s.l.Error(err, "failed to get banks")
		return nil, err
	}
	return rs.Data, nil
}
