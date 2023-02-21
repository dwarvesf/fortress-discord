package view

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/changelog"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/digest"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/done"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/errors"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/help"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/issue"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/memo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/project"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/techradar"
)

type View struct {
	subView subView
}

type subView struct {
	Earn       earn.EarnViewer
	Help       help.HelpViewer
	Error      errors.ErrorViewer
	TechRadar  techradar.TechRadarViewer
	Subscriber subscriber.SubscriberViewer
	Hiring     hiring.HiringViewer
	Event      event.EventViewer
	Staff      staff.StaffViewer
	Project    project.ProjectViewer
	Digest     digest.DigestViewer
	Memo       memo.MemoViewer
	Done       done.DoneViewer
	Issue      issue.IssueViewer
	Changelog  changelog.ChangelogViewer
}

func New(ses *discordgo.Session) Viewer {
	return &View{
		subView: subView{
			Earn:       earn.New(ses),
			Help:       help.New(ses),
			Error:      errors.New(ses),
			TechRadar:  techradar.New(ses),
			Subscriber: subscriber.New(ses),
			Hiring:     hiring.New(ses),
			Event:      event.New(ses),
			Staff:      staff.New(ses),
			Project:    project.New(ses),
			Digest:     digest.New(ses),
			Memo:       memo.New(ses),
			Done:       done.New(ses),
			Issue:      issue.New(ses),
			Changelog:  changelog.New(ses),
		},
	}
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

func (v *View) Done() done.DoneViewer {
	return v.subView.Done
}

func (v *View) Issue() issue.IssueViewer {
	return v.subView.Issue
}

func (v *View) Changelog() changelog.ChangelogViewer {
	return v.subView.Changelog
}
