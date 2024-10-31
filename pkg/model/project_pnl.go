package model

type ProjectPnL struct {
	Name               string  `json:"name"`
	Code               string  `json:"code"`
	EstimatedCost      float64 `json:"estimatedCost"`
	EstimatedRevenue   float64 `json:"estimatedRevenue"`
	RevenueToCostRatio float64 `json:"revenueToCostRatio"`
}
