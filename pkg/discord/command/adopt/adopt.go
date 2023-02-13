package adopt

import (
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Adopt struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) AdoptCommander {
	return &Adopt{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Adopt) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.TechRadar().GetList(constant.TechRadarRingAdopt, nil)
	if err != nil {
		t.L.Error(err, "can't get list of Adopt techradar")
		return err
	}

	// 2. render
	return t.view.TechRadar().ListAdopt(message, data)
}
