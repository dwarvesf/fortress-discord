package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/techradar"
)

type Servicer interface {
	Earn() earn.EarnServicer
	TechRadar() techradar.TechRadarServicer
	Subscriber() subscriber.SubscriberServicer
	Hiring() hiring.HiringServicer
	Event() event.EventServicer
	Staff() staff.StaffServicer
}
