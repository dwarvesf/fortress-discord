package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Memo struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) MemoServicer {
	return &Memo{
		adapter: adapter,
		l:       l,
	}
}

func (e *Memo) GetMemos() ([]*model.Memo, error) {
	// get response from fortress
	adapterMemos, err := e.adapter.Fortress().GetMemos()
	if err != nil {
		e.l.Error(err, "can't get open memo from fortress")
		return nil, err
	}

	// normalized into in-app model
	memos := adapterMemos.Data

	return memos, nil
}
