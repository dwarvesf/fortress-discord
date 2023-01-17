package service

import "github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"

type Service struct {
	subService subService
}

type subService struct {
	Earn earn.EarnServicer
}

func New() Servicer {
	return &Service{
		subService: subService{
			Earn: earn.New(),
		},
	}
}

func (s *Service) Earn() earn.EarnServicer {
	return s.subService.Earn
}
