package techradar

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TechRadarServicer interface {
	GetList(filterType string, query *string) ([]*model.TechRadarTopic, error)
}
