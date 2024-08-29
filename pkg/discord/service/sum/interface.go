package sum

import "github.com/dwarvesf/fortress-discord/pkg/model"

type SumServicer interface {
	SummarizeArticle(template, url string) (*model.Sum, error)
}
