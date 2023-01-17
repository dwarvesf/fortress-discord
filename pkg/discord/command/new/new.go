package new

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type NewCommand struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) NewCommander {
	return &NewCommand{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *NewCommand) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Subscriber().GetList()
	if err != nil {
		t.L.Error(err, "can't get list of subscriber")
		return err
	}

	// 2. render
	return t.view.Subscriber().ListNew(message, data)
}
