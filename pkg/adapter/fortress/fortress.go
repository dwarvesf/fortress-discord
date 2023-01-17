package fortress

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Fortress struct {
	Url string
}

func New(url string) FortressAdapter {
	return &Fortress{
		Url: url,
	}
}

func (f *Fortress) GetCommunityEarn() (earns *model.AdapterEarn, err error) {
	resp, err := http.Get(f.Url + "/api/v1/earn")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid call, code %v", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&earns); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return earns, nil
}
