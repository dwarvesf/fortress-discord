package view

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/errors"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/help"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/subscriber"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/techradar"
)

type Viewer interface {
	Earn() earn.EarnViewer
	Help() help.HelpViewer
	Error() errors.ErrorViewer
	TechRadar() techradar.TechRadarViewer
	Subscriber() subscriber.SubscriberViewer
	Hiring() hiring.HiringViewer
	Event() event.EventViewer
}
