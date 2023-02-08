package treasury

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Treasury struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) TreasuryServicer {
	return &Treasury{
		adapter: adapter,
		l:       l,
	}
}

func (t *Treasury) SendTip(tip *model.Tip) error {
	// check treasury fund

	// send request to mochi
	err := t.adapter.Mochi().SendTip(tip)
	if err != nil {
		t.l.Error(err, "can't send tip")
		return err
	}

	return nil
}
