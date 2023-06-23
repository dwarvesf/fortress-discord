package brainery

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Brainery struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) Commander {
	return &Brainery{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

type Commander interface {
	base.TextCommander

	Post(message *model.DiscordMessage) error
	Report(message *model.DiscordMessage) error
}
