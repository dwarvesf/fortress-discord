package staff

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type StaffViewer interface {
	ListDemands(original *model.DiscordMessage, subs []*model.StaffingDemand) error
	Help(original *model.DiscordMessage) error
}
