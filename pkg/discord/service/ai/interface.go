package ai

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type AIServicer interface {
	ProcessText(input string) (*model.AIResponse, error)
	ProcessTextWithN8N(input string) (*model.AIResponse, error)
}
