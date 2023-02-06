package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Memo struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) MemoCommander {
	return &Memo{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Memo) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Memo().GetMemos()
	if err != nil {
		t.L.Error(err, "can't get list of Memo")
		return err
	}

	// 2. render
	return t.view.Memo().List(message, data)
}
