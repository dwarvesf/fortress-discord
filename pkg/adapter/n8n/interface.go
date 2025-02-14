package n8n

import "github.com/dwarvesf/fortress-discord/pkg/model"

type N8nAdapter interface {
	ForwardPromptText(prompt, authorName, authorId, authorRoleId string) (*model.N8NEmbedResponse, error)
}
