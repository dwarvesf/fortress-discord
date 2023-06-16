package brainery

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Service interface {
	Post(input *model.Brainery) (*model.Brainery, error)
}
