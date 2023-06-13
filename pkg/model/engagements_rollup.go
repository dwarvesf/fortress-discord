package model

type EngagementsRollupRecord struct {
	DiscordUserID string `json:"discordUserID"`
	LastMessageID string `json:"lastMessageID"`
	ChannelID     string `json:"channelID"`
	CategoryID    string `json:"categoryID"`
	MessageCount  int    `json:"messageCount"`
	ReactionCount int    `json:"reactionCount"`
}
