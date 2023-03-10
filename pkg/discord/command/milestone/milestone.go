package milestone

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Milestone struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) MilestoneCommander {
	return &Milestone{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Milestone) ListMilestones(message *model.DiscordMessage) error {
	// this case we get all
	q := ""
	if len(message.ContentArgs) == 2 {
		q = message.ContentArgs[1]
	}

	// 1. get data from service
	data, err := t.svc.Project().GetListMilestone(q)
	if err != nil {
		t.L.Error(err, "can't get list of Milestones Project")
		return err
	}

	// 2. render
	if len(data) == 0 {
		return t.view.Project().EmptyMilestones(message)
	}

	return t.view.Project().ListMilestones(message, data)

}
