package view

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/errors"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/help"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/hiring"
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
