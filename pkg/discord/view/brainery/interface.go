package brainery

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Viewer interface {
	Help(message *model.DiscordMessage) error
	Post(original *model.DiscordMessage, content *model.Brainery, channelID string) error
	Report(original *model.DiscordMessage, view string, braineryMetric *model.BraineryMetric, channelID string) error
}
