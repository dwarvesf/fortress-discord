package fortress

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (f *Fortress) GetDiscordResearchTopics(page, size string) (data *model.DiscordResearchTopicResponse, err error) {
	req, err := f.makeReq(fmt.Sprintf("/api/v1/discords/research-topics?page=%s&size=%s", page, size), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[GetMemoOpenPullRequest] invalid call, code %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("[GetMemoOpenPullRequest] invalid decoded, error %v", err.Error())
	}

	return data, nil
}
