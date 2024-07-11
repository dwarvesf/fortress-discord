package researchtopic

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ResearchTopicServicer interface {
	GetDiscordResearchTopics(page, size string) (*model.DiscordResearchTopicResponse, error)
}
