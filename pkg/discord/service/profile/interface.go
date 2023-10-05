package profile

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Profile struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) Service {
	return &Profile{
		adapter: adapter,
		l:       l,
	}
}

type Service interface {
	Get(id string) (*model.Employee, error)
}
