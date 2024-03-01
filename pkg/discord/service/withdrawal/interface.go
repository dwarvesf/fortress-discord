package withdrawal

import "github.com/dwarvesf/fortress-discord/pkg/model"

// Servicer is the interface for withdraw service
type Servicer interface {
	CheckWithdrawCondition(discordID string) (*model.CheckWithdrawCondition, error)
	GetBanks(id, bin, swiftCode string) ([]model.Bank, error)
}
