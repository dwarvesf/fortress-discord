package fortress

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func (f *Fortress) GetTechRadar(ringFilter string, q *string) (techRadars *model.AdapterTechRadar, err error) {
	url := f.Url + "/api/v1/tech-radar?"
	if q != nil {
		url += "&name=" + *q
	}
	if ringFilter != "" {
		url += "&ring=" + ringFilter
	}
	resp, err := http.Get(url)
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
	resp, err := http.Get(f.Url + "/api/v1/audiences")
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
	resp, err := http.Get(f.Url + "/api/v1/hiring-positions")
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
	resp, err := http.Get(f.Url + "/api/v1/events")
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
	resp, err := http.Get(f.Url + "/api/v1/staffing-demands")
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
	resp, err := http.Get(f.Url + "/api/v1/projects/milestones?project_name=" + q)
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
	resp, err := http.Get(f.Url + "/api/v1/digests")
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
	resp, err := http.Get(f.Url + "/api/v1/updates")
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
	resp, err := http.Get(f.Url + "/api/v1/memos")
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
	resp, err := http.Get(f.Url + "/api/v1/issues")
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
	resp, err := http.Post(f.Url+"/api/v1/tech-radar", "application/json", bytes.NewBuffer(jsonValue))
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
