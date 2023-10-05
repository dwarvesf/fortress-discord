package view

import (
	"github.com/bwmarrin/discordgo"

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
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/profile"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/project"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/sum"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/techradar"
)

type View struct {
	subView subView
}

type subView struct {
	Brainery        brainery.Viewer
	Changelog       changelog.ChangelogViewer
	DeliveryMetrics deliverymetrics.DeliveryMetricsViewer
	Digest          digest.DigestViewer
	Done            done.DoneViewer
	Earn            earn.EarnViewer
	Error           errors.ErrorViewer
	Event           event.EventViewer
	Help            help.HelpViewer
	Hiring          hiring.HiringViewer
	Icy             icy.IcyViewer
	Issue           issue.IssueViewer
	Memo            memo.MemoViewer
	Profile         profile.Viewer
	Project         project.ProjectViewer
	Staff           staff.StaffViewer
	Subscriber      subscriber.SubscriberViewer
	Sum             sum.SumViewer
	TechRadar       techradar.TechRadarViewer
}

func New(ses *discordgo.Session) Viewer {
	return &View{
		subView: subView{
			Brainery:        brainery.New(ses),
			Changelog:       changelog.New(ses),
			DeliveryMetrics: deliverymetrics.New(ses),
			Digest:          digest.New(ses),
			Done:            done.New(ses),
			Earn:            earn.New(ses),
			Error:           errors.New(ses),
			Event:           event.New(ses),
			Help:            help.New(ses),
			Hiring:          hiring.New(ses),
			Icy:             icy.New(ses),
			Issue:           issue.New(ses),
			Memo:            memo.New(ses),
			Profile:         profile.New(ses),
			Project:         project.New(ses),
			Staff:           staff.New(ses),
			Subscriber:      subscriber.New(ses),
			Sum:             sum.New(ses),
			TechRadar:       techradar.New(ses),
		},
	}
}

func (v *View) Icy() icy.IcyViewer {
	return v.subView.Icy
}

func (v *View) Sum() sum.SumViewer {
	return v.subView.Sum
}

func (v *View) Earn() earn.EarnViewer {
	return v.subView.Earn
}

func (v *View) TechRadar() techradar.TechRadarViewer {
	return v.subView.TechRadar
}

func (v *View) Help() help.HelpViewer {
	return v.subView.Help
}

func (v *View) Subscriber() subscriber.SubscriberViewer {
	return v.subView.Subscriber
}

func (v *View) Error() errors.ErrorViewer {
	return v.subView.Error
}

func (v *View) Hiring() hiring.HiringViewer {
	return v.subView.Hiring
}

func (v *View) Event() event.EventViewer {
	return v.subView.Event
}

func (v *View) Staff() staff.StaffViewer {
	return v.subView.Staff
}

func (v *View) Project() project.ProjectViewer {
	return v.subView.Project
}

func (v *View) Digest() digest.DigestViewer {
	return v.subView.Digest
}

func (v *View) Memo() memo.MemoViewer {
	return v.subView.Memo
}

func (v *View) DeliveryMetrics() deliverymetrics.DeliveryMetricsViewer {
	return v.subView.DeliveryMetrics
}

func (v *View) Done() done.DoneViewer {
	return v.subView.Done
}

func (v *View) Issue() issue.IssueViewer {
	return v.subView.Issue
}

func (v *View) Changelog() changelog.ChangelogViewer {
	return v.subView.Changelog
}

func (v *View) Brainery() brainery.Viewer {
	return v.subView.Brainery
}

func (v *View) Profile() profile.Viewer {
	return v.subView.Profile
}
