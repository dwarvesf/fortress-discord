package digest

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Digest struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) DigestCommander {
	return &Digest{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Digest) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Digest().GetInteralUpdates()
	if err != nil {
		t.L.Error(err, "can't get list of Digest")
		return err
	}

	// 2. render
	return t.view.Digest().ListInternal(message, data)
}
