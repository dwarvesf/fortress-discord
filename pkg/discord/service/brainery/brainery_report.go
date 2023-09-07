package brainery

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Brainery) Report(view string, date string) (*model.BraineryMetric, error) {
	data, err := e.adapter.Fortress().GetBraineryReport(view, date)
	if err != nil {
		return nil, err
	}

	return data, nil
}
