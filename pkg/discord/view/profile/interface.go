package profile

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Viewer interface {
	Help(message *model.DiscordMessage) error
	Render(original *model.DiscordMessage, employees []model.Employee) error
	List(original *model.DiscordMessage, employees []model.Employee) error
	Details(original *model.DiscordMessage, employee *model.Employee) error
}
