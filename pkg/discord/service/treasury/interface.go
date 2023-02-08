package treasury

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TreasuryServicer interface {
	SendTip(tip *model.Tip) error
}
