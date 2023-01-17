package model

// AdapterEarn is a struct response from adapter, before process to in-app model
type AdapterEarn struct {
	Data    []*Earn
	Message string
}

// Earn is in-app model, after process from adapters
type Earn struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Reward int64    `json:"reward"`
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
}
