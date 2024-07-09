package news

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func (c svc) Reddit(subreddit string) ([]reddit.Post, []reddit.Post, error) {
	return c.adapter.Reddit().FetchNewsBySubreddit(context.Background(), subreddit)
}
