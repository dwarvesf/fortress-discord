package model

// AdapterSalaryAdvance is a struct response from adapter, before process to in-app model
type AdapterSalaryAdvance struct {
	Data    SalaryAdvance `json:"data"`
	Message string        `json:"message"`
}

// SalaryAdvance is in-app model, after process from adapters
type SalaryAdvance struct {
	AmountICY       string `json:"amountICY"`
	AmountUSD       string `json:"amountUSD"`
	TransactionID   string `json:"transactionID"`
	TransactionHash string `json:"transactionHash"`
}

type AdapterCheckSalaryAdvance struct {
	Data    CheckSalaryAdvance `json:"data"`
	Message string             `json:"message"`
}

type CheckSalaryAdvance struct {
	AmountICY string `json:"amountICY"`
	AmountUSD string `json:"amountUSD"`
}
