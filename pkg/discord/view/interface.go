package view

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/changelog"
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
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/project"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/sum"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/techradar"
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
	Digest() digest.DigestViewer
	Memo() memo.MemoViewer
	Done() done.DoneViewer
	Issue() issue.IssueViewer
	Changelog() changelog.ChangelogViewer
}
