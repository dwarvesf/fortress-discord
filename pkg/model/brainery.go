package model

import "time"

// Brainery is in-app model, after process from adapters
type Brainery struct {
	Title       string
	URL         string
	Author      string
	Description string
	Reward      string
	PublishDate *time.Time
	Tags        string
	Github      string
	DiscordID   string
}
