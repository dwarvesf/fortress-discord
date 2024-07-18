package researchtopic

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ResearchTopicServicer interface {
	GetDiscordResearchTopics(timeRange string) (*model.DiscordResearchTopicResponse, error)
}
