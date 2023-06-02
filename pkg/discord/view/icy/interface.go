package icy

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IcyViewer interface {
	List(original *model.DiscordMessage, earns []*model.Icy) error
	Help() error
}
