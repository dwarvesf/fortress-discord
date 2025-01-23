package project

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) GetProjectPnLs() ([]model.ProjectPnL, error) {
	pnls, err := e.adapter.IR().GetProjectPnLs()
	if err != nil {
		e.l.Error(err, "can't get project pnl")
		return nil, err
	}

	return pnls, nil
}
