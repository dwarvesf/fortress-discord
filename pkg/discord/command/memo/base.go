package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Memo) Prefix() []string {
	return []string{"memo", "memos"}
}

// Execute is where we handle logic for each command
func (e *Memo) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list":

		return e.ListMemoLogs(message)
	case "sync":
		return e.Sync(message)
	case "pr":
		return e.ListMemoOpenPullRequest(message)
	default:
		return e.ListByDiscordID(message)
	}
}

func (e *Memo) Name() string {
	return "Memo Command"
}

func (e *Memo) Help(message *model.DiscordMessage) error {
	return nil
}

// DefaultCommand handles the default command
func (e *Memo) DefaultCommand(message *model.DiscordMessage) error {
	return e.ListMemoLogs(message)
}

func (e *Memo) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
