package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProjectCommissionModel struct {
	Beneficiary    BasicEmployeeInfo       `json:"beneficiary"`
	CommissionType string                  `json:"type"`
	CommissionRate decimal.Decimal         `json:"commissionRate"`
	Description    string                  `json:"description"`
	Sub            *ProjectCommissionModel `json:"sub"`
}

type ProjectCommissionModelsResponse struct {
	Data []ProjectCommissionModel `json:"data"`
}

type Project struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Code         string     `json:"code"`
	Status       string     `json:"status"`
	ArtifactLink string     `json:"artifactLink"`
	SourceLink   string     `json:"sourceLink"`
	DocLink      string     `json:"docLink"`
	StartDate    *time.Time `json:"startDate"`
	EndDate      *time.Time `json:"endDate"`
}

type ProjectListResponse struct {
	Data []Project `json:"data"`
}
