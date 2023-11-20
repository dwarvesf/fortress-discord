package salary

import "github.com/dwarvesf/fortress-discord/pkg/model"

// SalaryServicer is the interface for Icy service
type SalaryServicer interface {
	SalaryAdvance(discordID, amount string) (*model.SalaryAdvance, error)
	CheckAdvanceSalary(discordID string) (*model.CheckSalaryAdvance, error)
}
