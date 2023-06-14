package sum

import "github.com/dwarvesf/fortress-discord/pkg/model"

type SumServicer interface {
	SummarizeArticle(url string) (*model.Sum, error)
}
