package fortress

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// SyncMemoLogs sync memos logs
func (f *Fortress) SyncMemoLogs() (memos *model.MemoLogsResponse, err error) {
	req, err := f.makeReq("/api/v1/memos/sync", http.MethodPost, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errMsg ErrorMessage
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return nil, errors.New("[SyncMemo] invalid decoded, error " + err.Error())
		}
		return nil, errors.New("[SyncMemo] invalid call, code " + strconv.Itoa(resp.StatusCode) + " " + errMsg.Message)
	}

	if err := json.NewDecoder(resp.Body).Decode(&memos); err != nil {
		return nil, fmt.Errorf("[GetMemoLogs] invalid decoded, error %v", err.Error())
	}

	return memos, nil
}

// GetMemoLogs get memos logs
func (f *Fortress) GetMemoLogs(from, to *time.Time) (memos *model.MemoLogsResponse, err error) {
	params := url.Values{}
	if from != nil {
		params.Add("from", from.Format(time.RFC3339))
	}
	if to != nil {
		params.Add("to", to.Format(time.RFC3339))
	}

	req, err := f.makeReq("/api/v1/memos?"+params.Encode(), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[GetMemoLogs] invalid call, code %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&memos); err != nil {
		return nil, fmt.Errorf("[GetMemoLogs] invalid decoded, error %v", err.Error())
	}

	return memos, nil
}

// GetMemoOpenPullRequest get open pull request
func (f *Fortress) GetMemoOpenPullRequest() (memos *model.MemoPullRequestResponse, err error) {
	req, err := f.makeReq("/api/v1/memos/prs", http.MethodGet, nil)
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

	if err := json.NewDecoder(resp.Body).Decode(&memos); err != nil {
		return nil, fmt.Errorf("[GetMemoOpenPullRequest] invalid decoded, error %v", err.Error())
	}

	return memos, nil
}

// GetMemoLogsByDiscordID get memos by discordID
func (f *Fortress) GetMemoLogsByDiscordID(discordID string) (memos *model.MemoLogsByDiscordIDResponse, err error) {
	params := url.Values{}
	if discordID != "" {
		params.Add("discordID", discordID)
	}

	req, err := f.makeReq("/api/v1/memos/discords?"+params.Encode(), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[GetMemoLogs] invalid call, code %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&memos); err != nil {
		return nil, fmt.Errorf("[GetMemoLogs] invalid decoded, error %v", err.Error())
	}

	return memos, nil
}

// GetTopAuthors get top authors
func (f *Fortress) GetTopAuthors(limit, days int) (topAuthors *model.MemoTopAuthorsResponse, err error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("days", strconv.Itoa(days))

	req, err := f.makeReq("/api/v1/memos/top-authors?"+params.Encode(), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[GetTopAuthors] invalid call, code %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&topAuthors); err != nil {
		return nil, fmt.Errorf("[GetTopAuthors] invalid decoded, error %v", err.Error())
	}

	return topAuthors, nil
}
