package fortress

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (f *Fortress) GetDiscordResearchTopics(timeRange string) (data *model.DiscordResearchTopicResponse, err error) {
	// only get top 5 topic
	req, err := f.makeReq(fmt.Sprintf("/api/v1/discords/research-topics?days=%s&page=1&size=5", timeRange), http.MethodGet, nil)
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
