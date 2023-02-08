package mochi

import "github.com/dwarvesf/fortress-discord/pkg/model"

type MochiAdapter interface {
	SendTip(tip *model.Tip) error
}
