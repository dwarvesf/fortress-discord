package earn

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EarnViewer interface {
	List(channelId string, earns []*model.Earn) error
	Help() error
}
