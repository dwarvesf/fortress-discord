package fortress

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (f *Fortress) GetDeliveryMetricsWeeklyReportDiscordMsg() (msg *model.AdapterDeliveryMetricsReportMsg, err error) {
	req, err := f.makeReq("/api/v1/delivery-metrics/report/weekly/discord-msg", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return msg, nil
}

func (f *Fortress) GetDeliveryMetricsMonthlyReportDiscordMsg() (msg *model.AdapterDeliveryMetricsReportMsg, err error) {
	req, err := f.makeReq("/api/v1/delivery-metrics/report/monthly/discord-msg", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return msg, nil
}

func (f *Fortress) SyncDeliveryMetricsData() (err error) {
	req, err := f.makeReq("/cronjobs/sync-delivery-metrics", http.MethodPost, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid call, code %v", resp.StatusCode)
	}

	return nil
}
