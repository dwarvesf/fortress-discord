package dify

type DifyAdapter interface {
	SummarizeArticle(youtubeURL string) (content string, err error)
}
