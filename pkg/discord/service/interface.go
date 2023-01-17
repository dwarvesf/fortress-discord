package service

import "github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"

type Servicer interface {
	Earn() earn.EarnServicer
}
