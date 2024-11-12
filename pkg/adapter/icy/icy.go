package icy

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Icy struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) IcyAdapter {
	return &Icy{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (i *Icy) GetBTCTreasury() (*model.IcyWeb3BigIntResponse, error) {
	resp, err := i.httpClient.Get(fmt.Sprintf("%s/api/v1/oracle/treasury-btc", i.baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.IcyWeb3BigIntResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (i *Icy) GetIcyRate() (*model.IcyWeb3BigIntResponse, error) {
	resp, err := i.httpClient.Get(fmt.Sprintf("%s/api/v1/oracle/icy-btc-ratio", i.baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.IcyWeb3BigIntResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
