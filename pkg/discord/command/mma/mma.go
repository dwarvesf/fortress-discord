package mma

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *MMACmd) ExportTemplate(message *model.DiscordMessage) error {
	data, err := e.svc.MMA().GetEmployeeMMAScore()
	if err != nil {
		return e.view.Error().Raise(message, "Failed to get employee profile.")
	}

	if len(data) == 0 {
		return e.view.Error().Raise(message, "There is no employee data with MMA score.")
	}

	return e.view.MMA().ExportTemplate(message, data)
}
