package n8n

import (
	"bytes"
	"encoding/json"
	"html"
	"io"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
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
func (n *N8n) ForwardPromptText(input, authorName, authorId string) (*model.N8NEmbedResponse, error) {
	payload := map[string]string{
		"content":     input,
		"author_name": authorName,
		"author_id":   authorId,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", n.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(n.WebhookUsername, n.WebhookPassword)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unescape the HTML characters in the response body
	unescapedBody := html.UnescapeString(string(bodyBytes))

	var embedResponse model.N8NEmbedResponse
	if err := json.Unmarshal([]byte(unescapedBody), &embedResponse); err != nil {
		return nil, err
	}

	return &embedResponse, nil
}
