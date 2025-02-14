package ai

import (
	"fmt"

	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type AI struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) AIServicer {
	return &AI{
		adapter: adapter,
		l:       l,
	}
}

func (a *AI) ProcessText(input string) (*model.AIResponse, error) {
	response, err := a.adapter.Dify().ProcessAIText(input)
	if err != nil {
		fmt.Printf("failed to process AI text. Error: %v", err)
		return nil, err
	}

	return &model.AIResponse{
		Input:    input,
		Response: response,
	}, nil
}

func (a *AI) ProcessTextWithN8N(input, authorId, authorName, authorRoleId string) (*model.N8NEmbedResponse, error) {
	embedResponse, err := a.adapter.N8n().ForwardPromptText(input, authorName, authorId, authorRoleId)
	if err != nil {
		fmt.Printf("failed to forward AI text to N8N webhook. Error: %v", err)
		return nil, err
	}

	return embedResponse, nil
}
