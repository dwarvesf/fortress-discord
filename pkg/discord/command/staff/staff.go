package staff

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Staff struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) StaffCommander {
	return &Staff{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Staff) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Staff().GetStaffingDemand()
	if err != nil {
		t.L.Error(err, "can't get list of active Staff")
		return err
	}

	// 2. render
	return t.view.Staff().ListDemands(message, data)
}
