package mochi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Mochi struct {
	Url string
}

func New(url string) MochiAdapter {
	return &Mochi{
		Url: url,
	}
}

func (m *Mochi) SendTip(tip *model.Tip) error {
	data, err := json.Marshal(tip)
	if err != nil {
		return err
	}

	resp, err := http.Post(m.Url+"/api/v1/tip/transfer", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		return errors.New(string(body))
	}

	return nil
}
