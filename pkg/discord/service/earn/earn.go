package earn

import (
	"sort"

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
	earns := make([]*model.Earn, len(adapterEarns.Data))
	earns = adapterEarns.Data

	// order by icy reward
	if len(earns) > 0 {
		sort.Slice(earns, func(i, j int) bool {
			return earns[i].Reward > earns[j].Reward
		})
	}

	return earns, nil
}
