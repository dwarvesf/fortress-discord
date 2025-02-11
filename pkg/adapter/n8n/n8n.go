package n8n

import (
	"bytes"
	"encoding/json"
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
func (n *N8n) ForwardPromptText(prompt string) error {
	payload := map[string]string{
		"prompt": prompt,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(n.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
