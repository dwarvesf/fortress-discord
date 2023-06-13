package engagement

import "github.com/dwarvesf/fortress-discord/pkg/model"

type EngagementServicer interface {
	UpsertRollup(record *model.EngagementsRollupRecord) error
}
