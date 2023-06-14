package openai

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
}

func New(apiKey string) OpenAIAdapter {
	client := openai.NewClient(apiKey)
	return &OpenAI{
		client: client,
	}
}

// SummarizeArticle implements OpenAIAdapter.
func (o *OpenAI) SummarizeArticle(url string) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Summarize this artile: %s. The summary should be 125 words max.", url),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
