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

type BraineryMetricItem struct {
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	Reward      decimal.Decimal `json:"reward"`
	PublishedAt string          `json:"publishedAt"`
	DiscordID   string          `json:"discordID"`
}

type TopContributor struct {
	DiscordID string
	Count     int
	Ranking   int
}

type BraineryMetric struct {
	LatestPosts     []BraineryMetricItem `json:"latestPosts"`
	Tags            []string             `json:"tags"`
	Contributors    []BraineryMetricItem `json:"contributors"`
	NewContributors []BraineryMetricItem `json:"newContributors"`
	TopContributors []TopContributor     `json:"topContributors"`
}

type BraineryMetricResponse struct {
	Data BraineryMetric `json:"data"`
}

type BasicEmployeeInfo struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
}
