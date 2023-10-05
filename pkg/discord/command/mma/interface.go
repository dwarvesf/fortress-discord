package mma

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type MMACmd struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) Commander {
	return &MMACmd{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

type Commander interface {
	base.TextCommander
}
