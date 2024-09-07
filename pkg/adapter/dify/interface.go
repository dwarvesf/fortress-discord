package dify

type DifyAdapter interface {
	SummarizeArticle(template, url string) (content string, err error)
}
