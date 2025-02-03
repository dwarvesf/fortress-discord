package techradar

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TechRadarViewer interface {
	ListTrial(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	ListAdopt(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	ListAssess(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	ListHold(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	Search(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	SearchEmpty(original *model.DiscordMessage) error

	LogTopicSuccess(original *model.DiscordMessage, topicName string) error
	LogTopicFailed(original *model.DiscordMessage, err string) error
	Help(original *model.DiscordMessage) error
}
