package brainery

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Service interface {
	Post(in *PostInput) (*model.Brainery, error)
}
