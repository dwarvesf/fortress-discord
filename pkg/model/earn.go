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

// MemoEarn represents new memo entity that is represented in the memo list
type MemoEarn struct {
	Title    string   `json:"title"`
	Bounty   string   `json:"bounty"`
	Status   string   `json:"status"`
	PICs     []string `json:"pics"`
	Function string   `json:"function"`
	URL      string   `json:"url"`
}

type AdapterMemoEarn struct {
	Data    []MemoEarn
	Message string
}
