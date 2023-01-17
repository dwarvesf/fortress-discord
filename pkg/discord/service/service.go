package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type Service struct {
	subService subService
}

type subService struct {
	Earn earn.EarnServicer
}

func New(adapter adapter.IAdapter, l logger.Logger) Servicer {
	return &Service{
		subService: subService{
			Earn: earn.New(adapter, l),
		},
	}
}

func (s *Service) Earn() earn.EarnServicer {
	return s.subService.Earn
}
