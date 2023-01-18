package fortress

import "github.com/dwarvesf/fortress-discord/pkg/model"

type FortressAdapter interface {
	GetCommunityEarn() (earns *model.AdapterEarn, err error)

	GetTechRadar(ringFilter string) (techradars *model.AdapterTechRadar, err error)

	GetNewSubscribers() (subscribers *model.AdapterSubscriber, err error)

	GetOpenPositions() (positions *model.AdapterHiringPosition, err error)

	GetUpcomingEvents() (events *model.AdapterEvent, err error)
}
