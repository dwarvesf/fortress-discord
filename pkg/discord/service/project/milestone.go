package project

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Project struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) ProjectServicer {
	return &Project{
		adapter: adapter,
		l:       l,
	}
}

func (e *Project) GetListMilestone(q string) ([]*model.ProjectMilestone, error) {
	// get response from fortress
	adapterMilestones, err := e.adapter.Fortress().GetProjectMilestones(q)
	if err != nil {
		e.l.Error(err, "can't get subscriber")
		return nil, err
	}

	// normalized into in-app model
	milestones := adapterMilestones.Data

	return milestones, nil
}

func (e *Project) GetProjectList() ([]model.Project, error) {
	return nil, nil
}
