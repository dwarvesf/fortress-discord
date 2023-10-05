package mma

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MMA struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) Service {
	return &MMA{
		adapter: adapter,
		l:       l,
	}
}

type Service interface {
	GetEmployeeMMAScore() ([]model.EmployeeMMAScore, error)
}
