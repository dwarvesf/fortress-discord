package earn

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EarnServicer interface {
	GetActiveList() ([]*model.Earn, error)
}
