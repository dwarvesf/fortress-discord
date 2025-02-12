package n8n

type N8nAdapter interface {
	ForwardPromptText(prompt, authorName, authorId string) (content string, err error)
}
