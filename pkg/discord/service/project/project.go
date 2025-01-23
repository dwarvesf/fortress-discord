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

func (e *Project) GetProjectList(q string) ([]model.Project, error) {
	res, err := e.adapter.Fortress().GetProjects(q)
	if err != nil {
		e.l.Error(err, "can't get projects")
		return nil, err
	}

	return res.Data, nil
}

func (e *Project) GetProjectByCode(q string) ([]model.Project, error) {
	res, err := e.adapter.Fortress().GetProjects(q)
	if err != nil {
		e.l.Error(err, "can't get project by code")
		return nil, err
	}

	return res.Data, nil
}
