package model

import "time"

type DiscordEvent struct {
	ID          string    `json:"id"`
	Name        string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	ChannelID   string    `json:"channel_id"`
	CreatorID   string    `json:"creator_id"`
	MessageID   string    `json:"message_id"`
}
