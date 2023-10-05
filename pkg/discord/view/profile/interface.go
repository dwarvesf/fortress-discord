package profile

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Viewer interface {
	Help(message *model.DiscordMessage) error
	Get(original *model.DiscordMessage, employee *model.Employee) error
}
