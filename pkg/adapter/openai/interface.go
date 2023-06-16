package openai

type OpenAIAdapter interface {
	SummarizeArticle(url string) (string, error)
	SummarizeBraineryPost(url string) (string, error)
}
