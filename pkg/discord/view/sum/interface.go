package sum

import "github.com/dwarvesf/fortress-discord/pkg/model"

type SumViewer interface {
	Sum(original *model.DiscordMessage, summary *model.Sum) error
	Help() error
}
