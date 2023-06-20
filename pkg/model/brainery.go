package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Brainery is in-app model, after process from adapters
type Brainery struct {
	Title       string
	URL         string
	Author      string
	Description string
	Reward      string
	PublishedAt *time.Time
	Tags        string
	Github      string
	DiscordID   string
}

type CreateBraineryLogRequest struct {
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	GithubID    string          `json:"githubID"`
	DiscordID   string          `json:"discordID"`
	Tags        []string        `json:"tags"`
	PublishedAt string          `json:"publishedAt"`
	Reward      decimal.Decimal `json:"reward"`
}