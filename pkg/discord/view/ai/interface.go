package ai

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type AIViewer interface {
	SendResponse(message *model.DiscordMessage, response *model.AIResponse) error
	Help(message *model.DiscordMessage) error
}
