package ir

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IRAdapter interface {
	GetProjectPnLs() ([]model.ProjectPnL, error)
}
