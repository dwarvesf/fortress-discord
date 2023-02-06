package model

import "time"

// AdapterDigest is a struct response from adapter, before process to in-app model
type AdapterDigest struct {
	Data    []*Digest `json:"data"`
	Message string    `json:"message"`
}

// Digest is in-app model, after process from adapters
type Digest struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}
