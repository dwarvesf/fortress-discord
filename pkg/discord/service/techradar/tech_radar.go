package techradar

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type TechRadar struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) TechRadarServicer {
	return &TechRadar{
		adapter: adapter,
		l:       l,
	}
}

func (e *TechRadar) GetList(ringFilter string, q *string) ([]*model.TechRadarTopic, error) {
	// get response from fortress
	adapterTechRadars, err := e.adapter.Fortress().GetTechRadar(ringFilter, q)
	if err != nil {
		e.l.Error(err, "can't get tech radar")
		return nil, err
	}

	// normalized into in-app model
	techRadars := adapterTechRadars.Data

	return techRadars, nil
}

func (e *TechRadar) LogTopic(topicName string, discordId string) error {
	// get response from fortress
	err := e.adapter.Fortress().LogTechRadarTopic(topicName, discordId)
	if err != nil {
		e.l.Error(err, "can't log tech radar")
		return err
	}

	return nil
}
