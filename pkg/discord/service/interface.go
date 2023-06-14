package service

import (
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
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/sum"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/techradar"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/treasury"
)

type Servicer interface {
	Earn() earn.EarnServicer
	Icy() icy.IcyServicer
	Sum() sum.SumServicer
	TechRadar() techradar.TechRadarServicer
	Subscriber() subscriber.SubscriberServicer
	Hiring() hiring.HiringServicer
	Event() event.EventServicer
	Staff() staff.StaffServicer
	Project() project.ProjectServicer
	Digest() digest.DigestServicer
	Changelog() changelog.ChangelogServicer
	Memo() memo.MemoServicer
	Treasury() treasury.TreasuryServicer
	Issue() issue.IssueServicer
}
