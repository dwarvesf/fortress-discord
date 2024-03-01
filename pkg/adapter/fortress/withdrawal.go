package fortress

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (f *Fortress) CheckWithdrawCondition(discordID string) (rs *model.AdapterCheckWithdrawCondition, err error) {
	req, err := f.makeReq(fmt.Sprintf("/api/v1/discords/withdraw/check?discordID=%v", discordID), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid call, code %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return rs, nil
}
