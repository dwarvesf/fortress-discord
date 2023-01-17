package model

import "time"

type Subscriber struct {
	Id        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Source    []string  `json:"source"`
	CreatedAt time.Time `json:"creatd_at"`
}

type AdapterSubscriber struct {
	Data    []*Subscriber `json:"data"`
	Message string        `json:"message"`
}
