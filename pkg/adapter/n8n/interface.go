package n8n

type N8nAdapter interface {
	ForwardPromptText(prompt string) error
}
