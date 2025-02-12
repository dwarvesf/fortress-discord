package df

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DFViewer interface {
	SendResponse(message *model.DiscordMessage, response *model.AIResponse) error
	Help(message *model.DiscordMessage) error
}
