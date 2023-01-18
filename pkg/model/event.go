package model

import "time"

type Event struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Date struct {
		Time    *time.Time `json:"time"`
		HasTime bool       `json:"has_time"`
	} `json:"date"`
}

type AdapterEvent struct {
	Data    []*Event `json:"data"`
	Message string   `json:"message"`
}
