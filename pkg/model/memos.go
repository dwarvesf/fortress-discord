package model

import (
	"time"

	"github.com/shopspring/decimal"
)

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

type CreateMemoRequest struct {
	Data []CreateMemoItem `json:"data"`
}

type CreateMemoItem struct {
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	Author      []Author        `json:"author"`
	Tags        []string        `json:"tags"`
	PublishedAt string          `json:"publishedAt"`
	Reward      decimal.Decimal `json:"reward"`
}

type Author struct {
	GithubID  string `json:"githubID"`
	DiscordID string `json:"discordID"`
}

type MemoLog struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	Authors     []MemoLogAuthor `json:"authors"`
	Description string          `json:"description"`
	PublishedAt *time.Time      `json:"publishedAt"`
	Reward      decimal.Decimal `json:"reward"`
	Category    []string        `json:"category"`
}

// MemoLogAuthor is the author of the memo log
type MemoLogAuthor struct {
	EmployeeID string `json:"employeeID"`
	GithubID   string `json:"githubID"`
	DiscordID  string `json:"discordID"`
}

// MemoLogsResponse response for memo logs
type MemoLogsResponse struct {
	Data []MemoLog `json:"data"`
}

type MemoPullRequest struct {
	Number         int       `json:"number"`
	Title          string    `json:"title"`
	DiscordId      string    `json:"discord_id"`
	GithubUserName string    `json:"github_user_name"`
	Url            string    `json:"url"`
	Timestamp      time.Time `json:"timestamp"`
}

type MemoRepoWithPullRequest map[string][]MemoPullRequest

// MemoPullRequestResponse response for memo open pull request
type MemoPullRequestResponse struct {
	Data MemoRepoWithPullRequest `json:"data"`
}

// MemoLogsByDiscordIDResponse response for memo logs by discord id
type MemoLogsByDiscordIDResponse struct {
	Data MemoLogsByDiscordID `json:"data"`
}

// AuthorRanking is the rank of the discord account
type AuthorRanking struct {
	DiscordID       string `json:"discordID"`
	DiscordUsername string `json:"discordUsername"`
	MemoUsername    string `json:"memoUsername"`
	TotalMemos      int    `json:"totalMemos"`
	Rank            int    `json:"rank"`
}

// MemoLogsByDiscordID response for memo logs by discord id
type MemoLogsByDiscordID struct {
	MemoLogs []MemoLog     `json:"memoLogs"`
	Rank     AuthorRanking `json:"rank"`
}

// MemoTopAuthorsResponse response for memo top authors
type MemoTopAuthorsResponse struct {
	Data []AuthorRanking `json:"data"`
} // @name MemoTopAuthorsResponse
