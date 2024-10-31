package project

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ProjectViewer interface {
	ListMilestones(original *model.DiscordMessage, milestones []*model.ProjectMilestone) error
	EmptyMilestones(original *model.DiscordMessage) error
	MissingArgsMilestones(original *model.DiscordMessage) error
	CommissionModels(original *model.DiscordMessage, commissionModel []model.ProjectCommissionModel) error
	PnL(original *model.DiscordMessage, pnls []model.ProjectPnL) error
}
