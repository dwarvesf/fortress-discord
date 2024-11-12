package icy

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IcyAdapter interface {
	GetBTCTreasury() (*model.IcyWeb3BigIntResponse, error)
	GetIcyRate() (*model.IcyWeb3BigIntResponse, error)
}
