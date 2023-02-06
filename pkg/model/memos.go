package model

import "time"

// AdapterMemo is a struct response from adapter, before process to in-app model
type AdapterMemo struct {
	Data    []*Memo `json:"data"`
	Message string  `json:"message"`
}

// Memo is in-app model, after process from adapters
type Memo struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}
