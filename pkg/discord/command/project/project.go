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
