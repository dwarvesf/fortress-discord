package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
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
}

func New(adapter adapter.IAdapter, l logger.Logger) Servicer {
	return &Service{
		subService: subService{
			Earn:       earn.New(adapter, l),
			TechRadar:  techradar.New(adapter, l),
			Subscriber: subscriber.New(adapter, l),
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
