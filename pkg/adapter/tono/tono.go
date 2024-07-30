package tono

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.New(string(body))
	}
	data := &model.ListGuildCommunityTransactionResponse{}
	err = json.Unmarshal(body, &data)
	return data, nil
}
