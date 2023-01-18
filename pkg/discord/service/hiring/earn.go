package hiring

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Hiring struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) HiringServicer {
	return &Hiring{
		adapter: adapter,
		l:       l,
	}
}

func (e *Hiring) GetOpenPositions() ([]*model.HiringPosition, error) {
	// get response from fortress
	adapterHiring, err := e.adapter.Fortress().GetOpenPositions()
	if err != nil {
		e.l.Error(err, "can't get open positions from fortress")
		return nil, err
	}

	// normalized into in-app model
	positions := adapterHiring.Data

	return positions, nil
}
