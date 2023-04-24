package fortress

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Fortress struct {
	Url    string
	ApiKey string
}

func New(url, apiKey string) FortressAdapter {
	return &Fortress{
		Url:    url,
		ApiKey: apiKey,
	}
}

func (f *Fortress) makeReq(subURL, method string, body io.Reader) (*http.Request, error) {
	url := f.Url + subURL
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "ApiKey "+f.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// GetChangelogs implements FortressAdapter
func (f *Fortress) GetChangelogs() (changelogs *model.ChangelogDigest, err error) {
	req, err := f.makeReq("/api/v1/notion/changelogs/projects/available", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&changelogs); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return changelogs, nil
}

func (f *Fortress) GetCommunityEarn() (earns *model.AdapterEarn, err error) {
	req, err := f.makeReq("/api/v1/notion/earn", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&earns); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return earns, nil
}

func (f *Fortress) GetTechRadar(ringFilter string, q *string) (techRadars *model.AdapterTechRadar, err error) {
	url := "/api/v1/notion/tech-radar?"
	if q != nil {
		url += "&name=" + *q
	}
	if ringFilter != "" {
		url += "&ring=" + ringFilter
	}

	req, err := f.makeReq(url, http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&techRadars); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return techRadars, nil
}

func (f *Fortress) GetNewSubscribers() (subscribers *model.AdapterSubscriber, err error) {
	req, err := f.makeReq("/api/v1/notion/audiences", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&subscribers); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return subscribers, nil
}

func (f *Fortress) GetOpenPositions() (posistions *model.AdapterHiringPosition, err error) {
	req, err := f.makeReq("/api/v1/notion/hiring-positions", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&posistions); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return posistions, nil
}

func (f *Fortress) GetUpcomingEvents() (events *model.AdapterEvent, err error) {
	req, err := f.makeReq("/api/v1/notion/events", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return events, nil
}

func (f *Fortress) GetStaffingDemands() (events *model.AdapterStaffingDemands, err error) {
	req, err := f.makeReq("/api/v1/notion/staffing-demands", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return events, nil
}

func (f *Fortress) GetProjectMilestones(q string) (milestone *model.AdapterProjectMilestone, err error) {
	req, err := f.makeReq("/api/v1/notion/projects/milestones?project_name="+q, http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&milestone); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return milestone, nil
}

func (f *Fortress) GetInternalDigest() (digest *model.AdapterDigest, err error) {
	req, err := f.makeReq("/api/v1/notion/digests", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&digest); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return digest, nil
}

func (f *Fortress) GetExternalDigest() (digest *model.AdapterDigest, err error) {
	req, err := f.makeReq("/api/v1/notion/updates", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&digest); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return digest, nil
}

func (f *Fortress) GetMemos() (memos *model.AdapterMemo, err error) {
	req, err := f.makeReq("/api/v1/notion/memos", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&memos); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return memos, nil
}

func (f *Fortress) GetActiveIssues() (issues *model.AdapterIssue, err error) {
	req, err := f.makeReq("/api/v1/notion/issues", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return issues, nil
}

func (f *Fortress) LogTechRadarTopic(topicName string, discordId string) error {
	type RadarTopic struct {
		Name      string `json:"name"`
		DiscordId string `json:"discord_id"`
	}

	// post to fortress
	radarTopic := RadarTopic{
		Name:      topicName,
		DiscordId: discordId,
	}

	jsonValue, _ := json.Marshal(radarTopic)
	req, err := f.makeReq("/api/v1/notion/tech-radar", http.MethodPost, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errMsg ErrorMessage
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return errors.New("invalid decoded, error " + err.Error())
		}
		return errors.New("invalid call, code " + strconv.Itoa(resp.StatusCode) + " " + errMsg.Message)
	}

	return nil
}

// SendChangelog implements FortressAdapter
func (f *Fortress) SendChangelog(c *model.Changelog) error {
	type SendChangelogReq struct {
		ProjectPageID string `json:"project_page_id"`
		IsPreview     bool   `json:"is_preview"`
		From          struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"from"`
	}

	// post to fortress
	bodyReq := SendChangelogReq{
		ProjectPageID: c.RowID,
		IsPreview:     false,
		From: struct {
			Email string "json:\"email\""
			Name  string "json:\"name\""
		}{
			Email: "team@d.foundation",
			Name:  "Team Dwarves",
		},
	}

	jsonValue, _ := json.Marshal(bodyReq)

	req, err := f.makeReq("/api/v1/notion/changelogs/project", http.MethodPost, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errMsg ErrorMessage
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return errors.New("invalid decoded, error " + err.Error())
		}
		return errors.New("invalid call, code " + strconv.Itoa(resp.StatusCode) + " " + errMsg.Message)
	}

	return nil
}
