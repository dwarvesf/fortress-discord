package trend

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TrendServicer interface {
	GetTrendingRepos(spokenLang string, programmingLang string, dateRange string) ([]*model.Repo, error)
}
