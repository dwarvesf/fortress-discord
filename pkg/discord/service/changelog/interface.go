package changelog

import "github.com/dwarvesf/fortress-discord/pkg/model"

type ChangelogServicer interface {
	GetListChangelogs() ([]*model.Changelog, error)
	SendChangelog(*model.Changelog) error
}
