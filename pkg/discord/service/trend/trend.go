package trend

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Trend struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) TrendServicer {
	return &Trend{
		adapter: adapter,
		l:       l,
	}
}

func (e *Trend) GetTrendingRepos(spokenLang string, programmingLang string, dateRange string) ([]*model.Repo, error) {
	// get response from fortress
	adapterTrends, err := e.adapter.Fortress().GetTrendingRepos(spokenLang, programmingLang, dateRange)
	if err != nil {
		e.l.Error(err, "can't get trending repos")
		return nil, err
	}

	// normalized into in-app model
	repos := adapterTrends.Data

	return repos, nil
}
