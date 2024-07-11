package view

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/brainery"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/changelog"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/deliverymetrics"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/digest"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/done"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/errors"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/help"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/icy"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/issue"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/memo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/mma"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/news"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/profile"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/project"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/salary"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/sum"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/techradar"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/topic"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/trend"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/withdrawal"
)

type Viewer interface {
	Earn() earn.EarnViewer
	Icy() icy.IcyViewer
	Sum() sum.SumViewer
	Help() help.HelpViewer
	Error() errors.ErrorViewer
	TechRadar() techradar.TechRadarViewer
	Subscriber() subscriber.SubscriberViewer
	Hiring() hiring.HiringViewer
	Staff() staff.StaffViewer
	Event() event.EventViewer
	Project() project.ProjectViewer
	DeliveryMetrics() deliverymetrics.DeliveryMetricsViewer
	Digest() digest.DigestViewer
	Memo() memo.MemoViewer
	Done() done.DoneViewer
	Issue() issue.IssueViewer
	Changelog() changelog.ChangelogViewer
	Brainery() brainery.Viewer
	Profile() profile.Viewer
	MMA() mma.Viewer
	Trend() trend.TrendViewer
	Salary() salary.Viewer
	Withdraw() withdrawal.Viewer
	News() news.Viewer
	Topic() topic.TopicViewer
}
