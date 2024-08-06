package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// Servicer is the interface for withdraw service
type Servicer interface {
	// Fetch calls to fortress to get news by platform and topic, using for general.
	Fetch(platform, topic string) ([]model.News, error)
}
