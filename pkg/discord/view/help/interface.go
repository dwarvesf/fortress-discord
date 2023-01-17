package help

import "github.com/dwarvesf/fortress-discord/pkg/model"

type HelpViewer interface {
	Help(message *model.DiscordMessage) error
}
