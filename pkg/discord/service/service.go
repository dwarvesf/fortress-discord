package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/changelog"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/digest"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/icy"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/issue"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/memo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/project"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/techradar"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/treasury"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type Service struct {
	subService subService
}

type subService struct {
	Earn       earn.EarnServicer
	Icy        icy.IcyServicer
	TechRadar  techradar.TechRadarServicer
	Subscriber subscriber.SubscriberServicer
	Hiring     hiring.HiringServicer
	Event      event.EventServicer
	Staff      staff.StaffServicer
	Project    project.ProjectServicer
	Digest     digest.DigestServicer
	Memo       memo.MemoServicer
	Treasury   treasury.TreasuryServicer
	Issue      issue.IssueServicer
	Changelog  changelog.ChangelogServicer
}

func New(adapter adapter.IAdapter, l logger.Logger) Servicer {
	return &Service{
		subService: subService{
			Earn:       earn.New(adapter, l),
			Icy:        icy.New(adapter, l),
			TechRadar:  techradar.New(adapter, l),
			Subscriber: subscriber.New(adapter, l),
			Hiring:     hiring.New(adapter, l),
			Event:      event.New(adapter, l),
			Staff:      staff.New(adapter, l),
			Project:    project.New(adapter, l),
			Digest:     digest.New(adapter, l),
			Memo:       memo.New(adapter, l),
			Treasury:   treasury.New(adapter, l),
			Issue:      issue.New(adapter, l),
			Changelog:  changelog.New(adapter, l),
		},
	}
}

// Icy implements Servicer.
func (s *Service) Icy() icy.IcyServicer {
	return s.subService.Icy
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

func (s *Service) Project() project.ProjectServicer {
	return s.subService.Project
}

func (s *Service) Digest() digest.DigestServicer {
	return s.subService.Digest
}

func (s *Service) Memo() memo.MemoServicer {
	return s.subService.Memo
}

func (s *Service) Treasury() treasury.TreasuryServicer {
	return s.subService.Treasury
}

func (s *Service) Issue() issue.IssueServicer {
	return s.subService.Issue
}

func (s *Service) Changelog() changelog.ChangelogServicer {
	return s.subService.Changelog
}
