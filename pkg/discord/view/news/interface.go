package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Viewer interface {
	Render(original *model.DiscordMessage, platform, topic string, posts []model.News) error
	Help(message *model.DiscordMessage) error
}
