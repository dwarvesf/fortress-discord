package earn

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Earn struct{}

func New() EarnServicer {
	return &Earn{}
}

func (e *Earn) GetActiveList() ([]*model.Earn, error) {
	return nil, nil
}
