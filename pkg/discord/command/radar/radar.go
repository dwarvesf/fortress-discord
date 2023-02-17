package radar

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Radar struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) RadarCommander {
	return &Radar{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (t *Radar) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.TechRadar().GetList(constant.TechRadarAll, nil)
	if err != nil {
		t.L.Error(err, "can't get list of radar techradar")
		return err
	}

	m := make(map[string][]*model.TechRadarTopic)

	for _, v := range data {
		m[v.Ring] = append(m[v.Ring], v)
	}

	for k, v := range m {
		if k == "Hold" {
			t.view.TechRadar().ListHold(message, v)
		}
		if k == "Assess" {
			t.view.TechRadar().ListAssess(message, v)
		}
		if k == "Trial" {
			t.view.TechRadar().ListTrial(message, v)
		}
		if k == "Adopt" {
			t.view.TechRadar().ListAdopt(message, v)
		}
	}

	// 2. render
	return nil
}

func (t *Radar) Log(message *model.DiscordMessage) error {
	// validate input
	if len(message.ContentArgs) <= 2 {
		// render error view
		return t.view.TechRadar().LogTopicFailed(message, "Missing topic name")
	}

	name := strings.Join(message.ContentArgs[2:], " ")
	name = strings.ReplaceAll(name, "\"", "")
	userID := message.Author.ID

	// send to fortress
	err := t.svc.TechRadar().LogTopic(name, userID)
	if err != nil {
		// render error view
		return t.view.TechRadar().LogTopicFailed(message, err.Error())
	}

	// render success view
	return t.view.TechRadar().LogTopicSuccess(message, name)
}
