package event

import (
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EventServicer interface {
	GetUpcomingEvents() ([]*model.Event, error)
	GetGuildScheduledEvents() ([]*model.DiscordEvent, error)
	CreateGuildScheduledEvent(*model.DiscordEvent) error
	SetSpeakers(message *model.DiscordMessage) error
	GetOgifStats(discordID string, after time.Time) (model.OgifStats, error)
	GetOgifLeaderboard(after time.Time, limit int) ([]model.OgifLeaderboardRecord, error)
}
