package news

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (c svc) Fetch(platform, topic string) ([]model.News, []model.News, error) {
	return c.adapter.Fortress().FetchNews(platform, topic)
}
