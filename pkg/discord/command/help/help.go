package help

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Help struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) HelpCommander {
	return &Help{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (h *Help) Name() string {
	return "Help Command"
}

func (h *Help) Help(message *model.DiscordMessage) error {
	return h.view.Help().Help(message)
}

func (h *Help) Prefix() []string {
	return []string{"help"}
}

func (h *Help) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return h.DefaultCommand(message)
	}

	// TODO: add sub help, for now only return help
	return h.DefaultCommand(message)
}

func (h *Help) DefaultCommand(message *model.DiscordMessage) error {
	return h.Help(message)
}
