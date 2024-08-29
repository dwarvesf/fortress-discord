package model

import "time"

// OgifStats contains list of ogif and some stats
type OgifStats struct {
	OgifList               []FortEventSpeaker `json:"ogifList"`
	UserAllTimeSpeaksCount int64              `json:"userAllTimeSpeaksCount"`
	UserAllTimeRank        int64              `json:"userAllTimeRank"`
	UserCurrentSpeaksCount int64              `json:"userCurrentSpeaksCount"`
	UserCurrentRank        int64              `json:"userCurrentRank"`
	TotalSpeakCount        int64              `json:"totalSpeakCount"`
	CurrentSpeakCount      int64              `json:"currentSpeakCount"`
}

// OgifStatsResponse return ogif stats response
type OgifStatsResponse struct {
	Data OgifStats `json:"data"`
} // @name OgifStatsResponse

// FortEventSpeaker struct
type FortEventSpeaker struct {
	EventID          string `json:"eventId"`
	DiscordAccountID string `json:"discordAccountId"`
	Topic            string `json:"topic,omitempty"`
	Event            *Event `json:"event"`
}

// FortEvent struct
type FortEvent struct {
	Date   time.Time `json:"date"`
	IsOver bool      `json:"isOver"`
}
