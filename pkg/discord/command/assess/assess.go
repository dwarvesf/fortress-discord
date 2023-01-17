package assess

import (
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Assess struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) AssessCommander {
	return &Assess{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Assess) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.TechRadar().GetList(constant.TechRadarRingAssess)
	if err != nil {
		t.L.Error(err, "can't get list of assess techradar")
		return err
	}

	// 2. render
	return t.view.TechRadar().ListAssess(message, data)
}
