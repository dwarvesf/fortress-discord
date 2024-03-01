package model

type WithdrawInput struct {
	BankName      string
	SwiftCode     string
	AccountNumber string

	ICYAmount  float64
	VNDAmount  float64
	ICYVNDRate float64
}

type AdapterCheckWithdrawCondition struct {
	Data CheckWithdrawCondition `json:"data"`
}

type CheckWithdrawCondition struct {
	ICYAmount  float64 `json:"icyAmount"`
	ICYVNDRate float64 `json:"icyVNDRate"`
	VNDAmount  float64 `json:"vndAmount"`
}
