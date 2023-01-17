package fortress

import "github.com/dwarvesf/fortress-discord/pkg/model"

type FortressAdapter interface {
	GetCommunityEarn() (earns *model.AdapterEarn, err error)
}
