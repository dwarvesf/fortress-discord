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

func (e *ProjectCmd) GetProjectPnL(message *model.DiscordMessage) error {
	// using curl to get project pnls
	pnls, err := e.svc.Project().GetProjectPnLs()
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get project PnLs.")
	}

	return e.view.Project().PnL(message, pnls)
}
