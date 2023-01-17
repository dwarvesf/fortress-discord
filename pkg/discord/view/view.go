package view

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/help"
)

type View struct {
	subView subView
}

type subView struct {
	Earn earn.EarnViewer
	Help help.HelpViewer
}

func New(ses *discordgo.Session) Viewer {
	return &View{
		subView: subView{
			Earn: earn.New(ses),
			Help: help.New(ses),
		},
	}
}

func (v *View) Earn() earn.EarnViewer {
	return v.subView.Earn
}

func (v *View) Help() help.HelpViewer {
	return v.subView.Help
}