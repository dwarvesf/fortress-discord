package view

import "github.com/dwarvesf/fortress-discord/pkg/discord/view/earn"

type Viewer interface {
	Earn() earn.EarnViewer
}
