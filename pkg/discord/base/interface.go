package base

import "github.com/dwarvesf/fortress-discord/pkg/model"

type Helper interface {
	Help()
}

type Executor interface {
	Execute(message *model.DiscordMessage) error
}

type DefaultCommander interface {
	DefaultCommand(message *model.DiscordMessage) error
}

type Prefixer interface {
	Prefix() []string
}

type Namer interface {
	Name() string
}

type TextCommander interface {
	Helper
	Executor
	DefaultCommander
	Prefixer
	Namer
}
