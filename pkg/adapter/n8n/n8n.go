package n8n

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type N8n struct {
	WebhookURL string
}

// New function return dify service
func New(WebhookURL string) *N8n {
	return &N8n{
		WebhookURL: WebhookURL,
	}
}

// ForwardPromptText forwards the prompt text from ?ai command to the N8n webhook
func (n *N8n) ForwardPromptText(input, authorName, authorId string) (content string, err error) {
	payload := map[string]string{
		"content":     input,
		"author_name": authorName,
		"author_id":   authorId,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(n.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read raw text response
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBytes), nil
}
