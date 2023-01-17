package hold

import (
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Hold struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) HoldCommander {
	return &Hold{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Hold) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.TechRadar().GetList(constant.TechRadarRingHold)
	if err != nil {
		t.L.Error(err, "can't get list of hold techradar")
		return err
	}

	// 2. render
	return t.view.TechRadar().ListHold(message, data)
}
