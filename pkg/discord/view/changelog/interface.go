package changelog

import "github.com/dwarvesf/fortress-discord/pkg/model"

// ChangelogViewer is an interface for changelog view
type ChangelogViewer interface {
	Changelog(message *model.DiscordMessage, data []*model.Changelog) error
	ChangelogSendSuccess(message *model.DiscordMessage, data *model.Changelog) error
	ChangelogSendFailed(message *model.DiscordMessage, data *model.Changelog) error
}
