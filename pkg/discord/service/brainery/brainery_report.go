package brainery

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Brainery) Report(view string) (*model.BraineryMetric, error) {
	data, err := e.adapter.Fortress().GetBraineryReport(view)
	if err != nil {
		return nil, err
	}

	return data, nil
}
