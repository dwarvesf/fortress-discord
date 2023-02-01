package staff

import "github.com/dwarvesf/fortress-discord/pkg/model"

type StaffServicer interface {
	GetStaffingDemand() ([]*model.StaffingDemand, error)
}
