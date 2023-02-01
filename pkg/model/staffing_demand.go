package model

type StaffingDemand struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Request string `json:"request"`
}

type AdapterStaffingDemands struct {
	Data    []*StaffingDemand `json:"data"`
	Message string            `json:"message"`
}
