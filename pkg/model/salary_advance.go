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

type AdapterSalaryAdvanceReport struct {
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
	Total int                  `json:"total"`
	Sort  string               `json:"sort"`
	Data  *SalaryAdvanceReport `json:"data"`
}

type SalaryAdvanceReport struct {
	TotalICY       int64                     `json:"totalICY"`
	TotalUSD       float64                   `json:"totalUSD"`
	SalaryAdvances []AggregatedSalaryAdvance `json:"salaryAdvances"`
}

type AggregatedSalaryAdvance struct {
	AmountICY       int64   `json:"amountICY"`
	AmountUSD       float64 `json:"amountUSD"`
	DiscordID       string  `json:"discordID"`
	DiscordUsername string  `json:"discordUsername"`
	EmployeeID      string  `json:"employeeID"`
}
