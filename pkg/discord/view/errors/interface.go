package errors

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ErrorViewer interface {
	// CommandNotFound send message to user when command not found
	CommandNotFound(m *model.DiscordMessage) error

	// NotHavePermission send message to user when user not have permission to execute command
	NotHavePermission(m *model.DiscordMessage, required []string) error
}
