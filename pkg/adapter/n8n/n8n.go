package n8n

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type N8n struct {
	WebhookURL      string
	WebhookUsername string
	WebhookPassword string
}

// New function return dify service
func New(WebhookURL, WebhookUsername, WebhookPassword string) *N8n {
	return &N8n{
		WebhookURL:      WebhookURL,
		WebhookUsername: WebhookUsername,
		WebhookPassword: WebhookPassword,
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

	req, err := http.NewRequest("POST", n.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(n.WebhookUsername, n.WebhookPassword)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBytes), nil
}
