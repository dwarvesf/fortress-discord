package earn

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EarnViewer interface {
	List(original *model.DiscordMessage, earns []*model.Earn) error
	ListMemoEarn(original *model.DiscordMessage, earns []model.MemoEarn) error
	Help(original *model.DiscordMessage) error
}
