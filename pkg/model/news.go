package model

import "time"

type News struct {
	Title      string    `json:"title"`
	URL        string    `json:"url"`
	Popularity int64     `json:"popularity"`
	CreatedAt  time.Time `json:"timestamp"`
}

type ListNews struct {
	Popular  []News `json:"popular"`
	Emerging []News `json:"emerging"`
}

type FetchNewsResponse struct {
	Data ListNews `json:"data"`
}
