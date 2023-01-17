package service

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/techradar"
)

type Servicer interface {
	Earn() earn.EarnServicer
	TechRadar() techradar.TechRadarServicer
}
