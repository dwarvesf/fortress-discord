package index

import (
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Index struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) IndexCommander {
	return &Index{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Index) Search(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.TechRadar().GetList(constant.TechRadarAll, &message.ContentArgs[1])
	if err != nil {
		t.L.Error(err, "can't get list of Index techradar")
		return err
	}

	// 2. render
	if len(data) == 0 {
		return t.view.TechRadar().SearchEmpty(message)
	}

	return t.view.TechRadar().Search(message, data)
}
