package n8n

import (
	"bytes"
	"encoding/json"
	"errors"
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
func (n *N8n) ForwardPromptText(prompt string) (content string, err error) {
	payload := map[string]string{
		"prompt": prompt,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(n.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	var results []struct {
		Output string `json:"output"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", errors.New("No response from N8N")
	}

	result := results[0]
	defer resp.Body.Close()

	return result.Output, nil
}
