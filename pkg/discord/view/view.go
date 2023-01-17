package view

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"
)

type View struct {
	subView subView
}

type subView struct {
	Earn earn.EarnViewer
}

func New(ses *discordgo.Session) Viewer {
	return &View{
		subView: subView{
			Earn: earn.New(ses),
		},
	}
}

func (v *View) Earn() earn.EarnViewer {
	return v.subView.Earn
}
