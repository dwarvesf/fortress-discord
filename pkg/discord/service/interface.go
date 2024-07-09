package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/brainery"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/changelog"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/deliverymetrics"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/digest"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/engagement"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/icy"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/issue"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/memo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/mma"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/news"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/profile"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/project"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/salary"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/sum"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/techradar"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/treasury"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/trend"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/withdrawal"
)

type Servicer interface {
	Brainery() brainery.Service
	Changelog() changelog.ChangelogServicer
	DeliveryMetrics() deliverymetrics.DeliveryMetricsServicer
	Digest() digest.DigestServicer
	Earn() earn.EarnServicer
	Engagement() engagement.EngagementServicer
	Event() event.EventServicer
	Hiring() hiring.HiringServicer
	Icy() icy.IcyServicer
	Salary() salary.SalaryServicer
	Issue() issue.IssueServicer
	Memo() memo.MemoServicer
	MMA() mma.Service
	Profile() profile.Service
	Project() project.ProjectServicer
	Staff() staff.StaffServicer
	Subscriber() subscriber.SubscriberServicer
	Sum() sum.SumServicer
	TechRadar() techradar.TechRadarServicer
	Treasury() treasury.TreasuryServicer
	Trend() trend.TrendServicer
	Withdrawal() withdrawal.Servicer
	News() news.Servicer
}
