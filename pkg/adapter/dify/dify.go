package dify

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Dify struct {
	BaseURL            string
	SummarizerAppToken string
	ProcessAIAppToken  string
}

// New function return dify service
func New(baseURL, summarizerAppToken, processAIAppToken string) *Dify {
	return &Dify{
		BaseURL:            baseURL,
		SummarizerAppToken: summarizerAppToken,
		ProcessAIAppToken:  processAIAppToken,
	}
}

// BaseEvent represents the common fields in the event
type BaseEvent struct {
	Event          string `json:"event,omitempty"`
	ConversationID string `json:"conversation_id,omitempty"`
	MessageID      string `json:"message_id,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	TaskID         string `json:"task_id,omitempty"`
	ID             string `json:"id,omitempty"`
	Position       int    `json:"position,omitempty"`
}

// AgentThought represents the specific fields for agent_thought events
type AgentThought struct {
	BaseEvent
	Thought      string      `json:"thought,omitempty"`
	Observation  string      `json:"observation,omitempty"`
	Tool         string      `json:"tool,omitempty"`
	ToolLabels   interface{} `json:"tool_labels,omitempty"`
	ToolInput    string      `json:"tool_input,omitempty"`
	MessageFiles interface{} `json:"message_files,omitempty"`
}

// AgentMessage represents the specific fields for agent_message events
type AgentMessage struct {
	BaseEvent
	Answer string `json:"answer,omitempty"`
}

func (d *Dify) SummarizeArticle(url string) (content string, err error) {
	// Define the URL and request body
	requestBody, err := json.Marshal(map[string]interface{}{
		"inputs":          map[string]interface{}{},
		"query":           url,
		"response_mode":   "streaming",
		"conversation_id": "",
		"user":            "fortress",
	})
	if err != nil {
		return "", err
	}

	return d.processDifyRequest(requestBody, d.SummarizerAppToken)
}

// ProcessAIText processes any text input using the Dify API
func (d *Dify) ProcessAIText(input string) (content string, err error) {
	// Define the URL and request body
	requestBody, err := json.Marshal(map[string]interface{}{
		"inputs":          map[string]interface{}{},
		"query":           input,
		"response_mode":   "streaming",
		"conversation_id": "",
		"user":            "fortress",
	})
	if err != nil {
		return "", err
	}

	return d.processDifyRequest(requestBody, d.ProcessAIAppToken)
}

// processDifyRequest handles the common logic for making requests to the Dify API
func (d *Dify) processDifyRequest(requestBody []byte, token string) (content string, err error) {
	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, d.BaseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// Set the required headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	thoughts, err := d.handleStreamingResponse(resp.Body)
	if err != nil {
		return "", err
	}

	// get the last event
	if len(thoughts) == 0 {
		return "", fmt.Errorf("no thought found")
	}

	for i := len(thoughts) - 1; i >= 0; i-- {
		if thoughts[i].Thought != "" {
			content = thoughts[i].Thought
			break
		}
	}

	return content, nil
}

// handleStreamingResponse processes the streaming response from the Dify API
func (d *Dify) handleStreamingResponse(body io.ReadCloser) ([]AgentThought, error) {
	var thoughts []AgentThought
	reader := bufio.NewReader(body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// Remove the "data: " prefix
		line = bytes.TrimPrefix(line, []byte("data: "))
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		// Parse the JSON event to a map to determine the event type
		var rawEvent map[string]interface{}
		err = json.Unmarshal(line, &rawEvent)
		if err != nil {
			fmt.Printf("Error unmarshal: %v\n", string(line))
			continue
		}

		eventType, ok := rawEvent["event"].(string)
		if !ok {
			fmt.Println("Error: event type is missing or not a string")
			continue
		}

		// Process specific event types
		switch eventType {
		case "agent_thought":
			var event AgentThought
			err = json.Unmarshal(line, &event)
			if err != nil {
				fmt.Printf("Error parsing agent_thought JSON: %v\n", err)
				continue
			}
			thoughts = append(thoughts, event)
		case "agent_message":
			// just ignore agent_message event
		default:
		}
	}

	return thoughts, nil
}
