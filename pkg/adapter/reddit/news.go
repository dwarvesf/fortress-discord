package reddit

import (
	"context"
	"sort"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func (a adapter) FetchNewsBySubreddit(ctx context.Context, sub string) ([]reddit.Post, []reddit.Post, error) {
	dayAgo := time.Now().Add(-24 * time.Hour)

	newPosts, redditResp, err := a.client.Subreddit.NewPosts(ctx, sub, &reddit.ListOptions{
		Limit: 50,
	})
	if err != nil {
		a.l.Errorf(err, "failed to fetch new posts with resp %v", redditResp)
		return nil, nil, err
	}

	newPostsMap := make(map[string]reddit.Post)
	for _, post := range newPosts {
		if post.Created.Before(dayAgo) {
			continue
		}

		newPostsMap[post.ID] = *post
	}

	risingPosts, redditResp, err := a.client.Subreddit.RisingPosts(ctx, sub, &reddit.ListOptions{
		Limit: 50,
	})
	if err != nil {
		a.l.Errorf(err, "failed to fetch rising posts with resp %v", redditResp)
		return nil, nil, err
	}

	risingPostsMap := make(map[string]reddit.Post)
	for _, post := range risingPosts {
		if post.Created.Before(dayAgo) {
			continue
		}

		risingPostsMap[post.ID] = *post
	}

	popularPosts := make([]reddit.Post, 0)
	emergingPosts := make([]reddit.Post, 0)
	for _, post := range newPostsMap {
		if _, ok := risingPostsMap[post.ID]; !ok {
			emergingPosts = append(emergingPosts, post)
			continue
		}

		popularPosts = append(popularPosts, post)
	}

	sort.Slice(popularPosts, func(i, j int) bool {
		return popularPosts[i].NumberOfComments > (popularPosts[j].NumberOfComments)
	})

	if len(popularPosts) > 10 {
		emergingPosts = append(emergingPosts, popularPosts[10:]...)
		popularPosts = popularPosts[:10]
	}

	sort.Slice(emergingPosts, func(i, j int) bool {
		return emergingPosts[i].Created.Time.After(emergingPosts[j].Created.Time)
	})

	if len(emergingPosts) > 10 {
		emergingPosts = emergingPosts[:10]
	}

	return popularPosts, emergingPosts, nil
}
