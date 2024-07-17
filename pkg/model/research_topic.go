package model

// DiscordResearchTopic represents discord research topic
type DiscordResearchTopic struct {
	Name              string
	URL               string
	MsgCount          int64
	SortedActiveUsers []DiscordTopicActiveUser
}

// DiscordTopicActiveUser represents active users who send most messages in topic
type DiscordTopicActiveUser struct {
	UserID   string
	MsgCount int64
}

type DiscordResearchTopicResponse struct {
	Page  int                    `json:"page"`
	Size  int                    `json:"size"`
	Sort  string                 `json:"sort"`
	Total int                    `json:"total"`
	Data  []DiscordResearchTopic `json:"data"`
}
