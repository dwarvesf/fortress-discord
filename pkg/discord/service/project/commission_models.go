package project

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) GetCommissionModels(projectID string) ([]model.ProjectCommissionModel, error) {
	// get response from fortress
	projectCommissionModels, err := e.adapter.Fortress().GetProjectCommissionModels(projectID)
	if err != nil {
		e.l.Error(err, "can't get subscriber")
		return nil, err
	}

	// normalized into in-app model
	commissionModels := projectCommissionModels

	return commissionModels, nil
}
