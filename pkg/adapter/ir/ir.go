package ir

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type ir struct {
	Url    string
	APIKey string
	Config *config.Config
}

func New(cfg *config.Config) IRAdapter {
	return &ir{
		Url:    cfg.Endpoint.IR.Url,
		APIKey: cfg.Endpoint.IR.APIKey,
		Config: cfg,
	}
}

func (m *ir) GetProjectPnLs() ([]model.ProjectPnL, error) {
	// Create a new request
	req, err := http.NewRequest("GET", m.Url+"/projects/pnl", nil)
	if err != nil {
		return nil, err
	}

	// Add the X-Api-Key header
	req.Header.Add("X-Api-Key", m.APIKey)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.New(string(body))
	}

	data := []model.ProjectPnL{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
