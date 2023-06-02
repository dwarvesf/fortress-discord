package model

// AdapterIcy is a struct response from adapter, before process to in-app model
type AdapterIcy struct {
	Data    []*Icy
	Message string
}

// Icy is in-app model, after process from adapters
type Icy struct {
	ID     string `json:"id"`
	Period string `json:"period"`
	Team   string `json:"team"`
	Amount string `json:"amount"`
}
