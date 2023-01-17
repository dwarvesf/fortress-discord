package earn

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EarnViewer interface {
	List(earns []*model.Earn) error
	Help() error
}
