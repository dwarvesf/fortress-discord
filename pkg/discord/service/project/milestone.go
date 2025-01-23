package project

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) GetListMilestone(q string) ([]*model.ProjectMilestone, error) {
	// get response from fortress
	adapterMilestones, err := e.adapter.Fortress().GetProjectMilestones(q)
	if err != nil {
		e.l.Error(err, "can't get milestones")
		return nil, err
	}

	// normalized into in-app model
	milestones := adapterMilestones.Data

	return milestones, nil
}
