package errors

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// ErrorViewer is an interface error view for some common errors
type ErrorViewer interface {
	// CommandNotFound send message to user when command not found
	CommandNotFound(m *model.DiscordMessage) error

	// CommandTemporarilyDisabled send message to user when command is disabled
	CommandTemporarilyDisabled(m *model.DiscordMessage) error

	// NotHavePermission send message to user when user not have permission to execute command
	NotHavePermission(m *model.DiscordMessage, required []string) error

	// Raise send message to user when error occur
	Raise(original *model.DiscordMessage, errorMessage string) error
}
