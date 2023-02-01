package staff

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Staff struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) StaffServicer {
	return &Staff{
		adapter: adapter,
		l:       l,
	}
}

func (e *Staff) GetStaffingDemand() ([]*model.StaffingDemand, error) {
	// get response from fortress
	adapterStaffing, err := e.adapter.Fortress().GetStaffingDemands()
	if err != nil {
		e.l.Error(err, "can't get open positions from fortress")
		return nil, err
	}

	// normalized into in-app model
	staffingDemands := adapterStaffing.Data

	return staffingDemands, nil
}
