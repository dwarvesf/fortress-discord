package project

import "github.com/dwarvesf/fortress-discord/pkg/model"

type ProjectServicer interface {
	GetListMilestone(q string) ([]*model.ProjectMilestone, error)
	GetCommissionModels(projectID string) ([]model.ProjectCommissionModel, error)
}
