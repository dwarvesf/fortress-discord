package tono

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Tono struct {
	Url    string
	Config *config.Config
}

func New(cfg *config.Config) TonoAdapter {
	return &Tono{
		Url:    cfg.Endpoint.Tono,
		Config: cfg,
	}
}

func (m *Tono) GetCommunityTransaction() (*model.ListGuildCommunityTransactionResponse, error) {
	resp, err := http.Get(m.Url + fmt.Sprintf("/api/v1/guilds/%s/community/transactions?interval=30", m.Config.Discord.ID.DwarvesGuild))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.New(string(body))
	}
	data := &model.ListGuildCommunityTransactionResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
