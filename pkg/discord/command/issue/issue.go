package issue

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Issue struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) IssueCommander {
	return &Issue{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Issue) List(message *model.DiscordMessage) error {
	// 1. get data from service
	// data, err := t.svc.TechRadar().GetList(constant.TechRadarRingIssue, nil)
	// if err != nil {
	// 	t.L.Error(err, "can't get list of Issue")
	// 	return err
	// }

	// // 2. render
	// return t.view.TechRadar().ListAdopt(message, data)
	return nil
}
