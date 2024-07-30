package tono

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Tono struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) TonoServicer {
	return &Tono{
		adapter: adapter,
		l:       l,
	}
}

func (t *Tono) GetCommunityTransaction() (*model.ListGuildCommunityTransaction, error) {
	data, err := t.adapter.Tono().GetCommunityTransaction()
	if err != nil {
		t.l.Error(err, "can't get tono community transaction")
		return nil, err
	}

	return &data.Data, nil
}
