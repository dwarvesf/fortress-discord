package model

import "time"

type News struct {
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	Popularity   int64     `json:"popularity"`
	CommentCount int64     `json:"comment_count"`
	Flag         int64     `json:"flag"`
	Description  string    `json:"description"`
	Tags         []string  `json:"tags"`
	CreatedAt    time.Time `json:"timestamp"`
}

type FetchNewsResponse struct {
	Data []News `json:"data"`
}
