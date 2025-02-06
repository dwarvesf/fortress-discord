package techradar

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TechRadarViewer interface {
	ListTrial(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	TrialHelp(original *model.DiscordMessage) error

	ListAdopt(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	AdoptHelp(original *model.DiscordMessage) error

	ListAssess(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	AssessHelp(original *model.DiscordMessage) error

	ListHold(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	HoldHelp(original *model.DiscordMessage) error

	Search(original *model.DiscordMessage, topics []*model.TechRadarTopic) error
	SearchEmpty(original *model.DiscordMessage) error
	IndexHelp(original *model.DiscordMessage) error

	LogTopicSuccess(original *model.DiscordMessage, topicName string) error
	LogTopicFailed(original *model.DiscordMessage, err string) error
}
