package earn

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Earn struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) EarnServicer {
	return &Earn{
		adapter: adapter,
		l:       l,
	}
}

func (e *Earn) GetActiveList() ([]*model.Earn, error) {
	// get response from fortress
	adapterEarns, err := e.adapter.Fortress().GetCommunityEarn()
	if err != nil {
		e.l.Error(err, "can't get community earn")
		return nil, err
	}

	// normalized into in-app model
	earns := adapterEarns.Data

	return earns, nil
}
