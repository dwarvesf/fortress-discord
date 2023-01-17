package view

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/errors"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/help"
)

type View struct {
	subView subView
}

type subView struct {
	Earn  earn.EarnViewer
	Help  help.HelpViewer
	Error errors.ErrorViewer
}

func New(ses *discordgo.Session) Viewer {
	return &View{
		subView: subView{
			Earn:  earn.New(ses),
			Help:  help.New(ses),
			Error: errors.New(ses),
		},
	}
}

func (v *View) Earn() earn.EarnViewer {
	return v.subView.Earn
}

func (v *View) Help() help.HelpViewer {
	return v.subView.Help
}

func (v *View) Error() errors.ErrorViewer {
	return v.subView.Error
}
