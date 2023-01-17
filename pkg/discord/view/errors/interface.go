package errors

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ErrorViewer interface {
	// CommandNotFound send message to user when command not found
	CommandNotFound(m *model.DiscordMessage) error
}
