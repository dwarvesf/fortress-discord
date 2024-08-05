package news

import "github.com/dwarvesf/fortress-discord/pkg/model"

func (c command) Fetch(msg *model.DiscordMessage, platform, tag string) error {
	logger := c.L.Field("func", "Lobsters")

	posts, err := c.svc.News().Fetch(platform, tag)
	if err != nil {
		logger.Error(err, "failed to fetch Golang news")
		return err
	}

	return c.view.News().Render(msg, platform, tag, posts)
}
