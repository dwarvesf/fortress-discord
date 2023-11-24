package salary

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Viewer interface {
	CompleteAdvanceSalary(original *model.DiscordMessage, salaryAdvance model.SalaryAdvance) error
	ErrorAdvanceSalary(original *model.DiscordMessage, err error) error
	EnterAmountAdvanceSalary(original *model.DiscordMessage, checkSalaryAdvance model.CheckSalaryAdvance) error
	Help() error
}
