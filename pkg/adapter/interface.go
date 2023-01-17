package adapter

import "github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"

type IAdapter interface {
	Fortress() fortress.FortressAdapter
}
