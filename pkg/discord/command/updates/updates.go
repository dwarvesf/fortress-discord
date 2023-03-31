package updates

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Updates struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) UpdatesCommander {
	return &Updates{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Updates) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Digest().GetExternalUpdates()
	if err != nil {
		t.L.Error(err, "can't get list of Updates")
		return err
	}

	// 2. render
	return t.view.Digest().ListExternal(message, data)
}

func (t *Updates) PreSend(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Digest().GetExternalUpdates()
	if err != nil {
		t.L.Error(err, "can't get list of Updates")
		return err
	}

	// 2. render
	return t.view.Digest().SendoutSelection(message, data)
}
