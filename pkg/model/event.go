package model

import "time"

type Event struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Date EventDate `json:"date"`
}

type EventDate struct {
	Time    *time.Time `json:"time"`
	HasTime bool       `json:"has_time"`
}

type AdapterEvent struct {
	Data    []*Event `json:"data"`
	Message string   `json:"message"`
}
