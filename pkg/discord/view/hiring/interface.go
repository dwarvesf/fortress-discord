package hiring

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type HiringViewer interface {
	OpeningList(original *model.DiscordMessage, subs []*model.HiringPosition) error
}
