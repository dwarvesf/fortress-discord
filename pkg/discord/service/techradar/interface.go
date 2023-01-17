package techradar

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TechRadarServicer interface {
	GetList(filterType string) ([]*model.TechRadarTopic, error)
}
