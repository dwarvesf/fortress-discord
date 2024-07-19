package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// Servicer is the interface for withdraw service
type Servicer interface {
	// TODO: update to use Fetch
	Reddit(subreddit string) ([]reddit.Post, []reddit.Post, error)
	// Fetch calls to fortress to get news by platform and topic, using for general.
	Fetch(platform, topic string) ([]model.News, []model.News, error)
}
