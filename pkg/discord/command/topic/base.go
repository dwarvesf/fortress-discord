package topic

import "github.com/dwarvesf/fortress-discord/pkg/model"

func (e *Topic) Prefix() []string {
	return []string{"topic", "topics"}
}

// Execute is where we handle logic for each command
func (e *Topic) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return e.DefaultCommand(message)
	}

	return nil
}

func (e *Topic) Name() string {
	return "Memo Command"
}

func (e *Topic) Help(message *model.DiscordMessage) error {
	return nil
}

// DefaultCommand handles the default command
func (e *Topic) DefaultCommand(message *model.DiscordMessage) error {
	return e.List(message)
}

func (e *Topic) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
