package hiring

import "github.com/dwarvesf/fortress-discord/pkg/model"

type HiringServicer interface {
	GetOpenPositions() ([]*model.HiringPosition, error)
}
