package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Adapter interface {
	FetchNewsBySubreddit(ctx context.Context, sub string) ([]reddit.Post, []reddit.Post, error)
}
