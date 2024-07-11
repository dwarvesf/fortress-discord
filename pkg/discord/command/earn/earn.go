package earn

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Earn struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) EarnCommander {
	return &Earn{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Earn) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Earn().ListMemoEarn()
	if err != nil {
		e.L.Error(err, "can't get list of active earn")
		return err
	}

	// 2. render
	return e.view.Earn().ListMemoEarn(message, data)
}
