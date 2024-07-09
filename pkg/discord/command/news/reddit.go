package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c command) Reddit(msg *model.DiscordMessage, subreddit string) error {
	logger := c.L.Field("func", "Reddit")

	popular, emerging, err := c.svc.News().Reddit(subreddit)
	if err != nil {
		logger.Error(err, "failed to fetch Golang news")
		return err
	}

	if len(popular) == 0 && len(emerging) == 0 {
		logger.Info("no new Golang news")
		return nil
	}

	return c.view.News().Reddit(msg, subreddit, popular, emerging)
}
