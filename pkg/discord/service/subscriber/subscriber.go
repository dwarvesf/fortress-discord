package subscriber

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Subsriber struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) SubscriberServicer {
	return &Subsriber{
		adapter: adapter,
		l:       l,
	}
}

func (e *Subsriber) GetList() ([]*model.Subscriber, error) {
	// get response from fortress
	adapterSubscribers, err := e.adapter.Fortress().GetNewSubscribers()
	if err != nil {
		e.l.Error(err, "can't get subscriber")
		return nil, err
	}

	// normalized into in-app model
	subscribers := adapterSubscribers.Data

	return subscribers, nil
}
