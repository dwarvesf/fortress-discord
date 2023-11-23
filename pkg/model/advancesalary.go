package model

// AdapterSalaryAdvance is a struct response from adapter, before process to in-app model
type AdapterSalaryAdvance struct {
	Data    SalaryAdvance `json:"data"`
	Message string        `json:"message"`
}

// SalaryAdvance is in-app model, after process from adapters
type SalaryAdvance struct {
	AmountIcy       string `json:"amount_icy"`
	AmountUSD       string `json:"amount_usd"`
	TransactionID   string `json:"transaction_id"`
	TransactionHash string `json:"transaction_hash"`
}

type AdapterCheckSalaryAdvance struct {
	Data    CheckSalaryAdvance `json:"data"`
	Message string             `json:"message"`
}

type CheckSalaryAdvance struct {
	AmountIcy string `json:"amount_icy"`
	AmountUSD string `json:"amount_usd"`
}
