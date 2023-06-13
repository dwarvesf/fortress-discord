package engagement

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Engagement struct {
	a adapter.IAdapter
	l logger.Logger
}

func New(a adapter.IAdapter, l logger.Logger) EngagementServicer {
	return &Engagement{
		a: a,
		l: l,
	}
}

func (e *Engagement) UpsertRollup(record *model.EngagementsRollupRecord) error {
	err := e.a.Fortress().UpsertRollupRecord(record)
	if err != nil {
		e.l.Error(err, "upsert rollup record error")
		return err
	}

	return nil
}
