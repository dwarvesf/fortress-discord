package mma

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Viewer interface {
	Help(message *model.DiscordMessage) error
	ExportTemplate(original *model.DiscordMessage, data []model.EmployeeMMAScore) error
}
