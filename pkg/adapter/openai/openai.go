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
					Role: openai.ChatMessageRoleSystem,
					Content: `You are able to accurately read news, books, and summaries.
					 you should be able to comprehend the content and provide a detailed understanding of the material.
					 Additionally, you should be able to identify key points and summarize the material in an organized manner.
					 You should also be able to recognize any errors or inconsistencies in the material and provide feedback on how to improve it.
					`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`Summary all keys takeaway from this article: %s. Format into max 3 concise bullet point`, url),
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
