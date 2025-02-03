package issue

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type IssueViewer interface {
	ListActive(original *model.DiscordMessage, subs []*model.Issue) error
	Help(original *model.DiscordMessage) error
}
