package project

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *ProjectCmd) GetProjectCommissionModels(message *model.DiscordMessage) error {
	projectID := message.ContentArgs[1]

	commissionModels, err := e.svc.Project().GetCommissionModels(projectID)
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get commission models.")
	}

	return e.view.Project().CommissionModels(message, commissionModels)
}

func (e *ProjectCmd) GetProjectList(message *model.DiscordMessage) error {
	status := "active"
	if len(message.ContentArgs) >= 3 {
		status = message.ContentArgs[2]
	}
	// using curl to get project pnls
	projs, err := e.svc.Project().GetProjectList(status)
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get project PnLs.")
	}

	return e.view.Project().List(message, projs)
}

func (e *ProjectCmd) GetProjectPnL(message *model.DiscordMessage) error {
	// using curl to get project pnls
	pnls, err := e.svc.Project().GetProjectPnLs()
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get project PnLs.")
	}

	return e.view.Project().PnL(message, pnls)
}
