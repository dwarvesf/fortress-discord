package model

import "time"

type DiscordEvent struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Date             time.Time `json:"date"`
	DiscordEventID   string    `json:"discord_event_id"`
	DiscordChannelID string    `json:"discord_channel_id"`
	DiscordCreatorID string    `json:"discord_creator_id"`
	DiscordMessageID string    `json:"discord_message_id"`
	IsOver           bool      `json:"is_over"`
}
