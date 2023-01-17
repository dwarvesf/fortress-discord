package techradar

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TechRadarViewer interface {
	ListTrial(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	ListAdopt(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	ListAssess(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	ListHold(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
}
