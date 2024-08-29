package ogif

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type command struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) Commander {
	return &command{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}
