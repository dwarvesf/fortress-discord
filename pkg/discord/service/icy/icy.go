package icy

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Icy struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) IcyServicer {
	return &Icy{
		adapter: adapter,
		l:       l,
	}
}

func (e *Icy) GetWeeklyDistribution() ([]*model.Icy, error) {
	// get response from fortress
	adapterIcys, err := e.adapter.Fortress().GetIcyWeeklyDistribution()
	if err != nil {
		e.l.Error(err, "can't get weekly icy distribution")
		return nil, err
	}

	// normalized into in-app model
	icys := adapterIcys.Data

	return icys, nil
}
