package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/techradar"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type Service struct {
	subService subService
}

type subService struct {
	Earn       earn.EarnServicer
	TechRadar  techradar.TechRadarServicer
	Subscriber subscriber.SubscriberServicer
	Hiring     hiring.HiringServicer
	Event      event.EventServicer
	Staff      staff.StaffServicer
}

func New(adapter adapter.IAdapter, l logger.Logger) Servicer {
	return &Service{
		subService: subService{
			Earn:       earn.New(adapter, l),
			TechRadar:  techradar.New(adapter, l),
			Subscriber: subscriber.New(adapter, l),
			Hiring:     hiring.New(adapter, l),
			Event:      event.New(adapter, l),
			Staff:      staff.New(adapter, l),
		},
	}
}

func (s *Service) Earn() earn.EarnServicer {
	return s.subService.Earn
}

func (s *Service) TechRadar() techradar.TechRadarServicer {
	return s.subService.TechRadar
}

func (s *Service) Subscriber() subscriber.SubscriberServicer {
	return s.subService.Subscriber
}

func (s *Service) Hiring() hiring.HiringServicer {
	return s.subService.Hiring
}

func (s *Service) Event() event.EventServicer {
	return s.subService.Event
}

func (s *Service) Staff() staff.StaffServicer {
	return s.subService.Staff
}
