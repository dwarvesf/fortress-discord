package adapter

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/adapter/mochi"
)

type IAdapter interface {
	Fortress() fortress.FortressAdapter
	Mochi() mochi.MochiAdapter
}
