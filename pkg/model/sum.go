package model

// AdapterSum is a struct response from adapter, before process to in-app model
type AdapterSum struct {
	Data    *Sum
	Message string
}

// Sum is in-app model, after process from adapters
type Sum struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Template string `json:"template"`
}
