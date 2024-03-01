package index

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (i *Index) Prefix() []string {
	return []string{"index"}
}

// Execute is where we handle logic for each command
func (i *Index) Execute(message *model.DiscordMessage) error {
	return i.DefaultCommand(message)
}

func (i *Index) Name() string {
	return "Home Command"
}

func (i *Index) Help(message *model.DiscordMessage) error {
	return nil
}

func (i *Index) DefaultCommand(message *model.DiscordMessage) error {
	return i.Search(message)
}

func (i *Index) PermissionCheck(message *model.DiscordMessage) (bool, []string) {
	return true, []string{}
}
