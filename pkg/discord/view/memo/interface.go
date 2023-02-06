package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MemoViewer interface {
	List(original *model.DiscordMessage, subs []*model.Memo) error
}
