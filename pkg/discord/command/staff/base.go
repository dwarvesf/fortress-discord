package staff

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (s *Staff) Prefix() []string {
	return []string{"staff"}
}

// Execute is where we handle logic for each command
func (s *Staff) Execute(message *model.DiscordMessage) error {
	// default command for only 1 args input from user, e.g `?earn`
	if len(message.ContentArgs) == 1 {
		return s.DefaultCommand(message)
	}

	// handle command for 2 args input from user, e.g `?earn list`
	switch message.ContentArgs[1] {
	case "list", "ls":
		return s.List(message)
	case "help", "h":
		return s.Help(message)
	default:
		return s.Help(message)
	}

	return nil
}

func (s *Staff) Name() string {
	return "Staff Command"
}

func (s *Staff) Help(message *model.DiscordMessage) error {
	return s.view.Staff().Help(message)
}

func (s *Staff) DefaultCommand(message *model.DiscordMessage) error {
	return s.Help(message)
}

func (s *Staff) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
