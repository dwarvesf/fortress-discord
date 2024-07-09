package news

import "github.com/vartanbeno/go-reddit/v2/reddit"

// Servicer is the interface for withdraw service
type Servicer interface {
	Reddit(subreddit string) ([]reddit.Post, []reddit.Post, error)
}
