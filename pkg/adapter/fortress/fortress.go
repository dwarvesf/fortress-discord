package fortress

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"

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

func (f *Fortress) GetIcyWeeklyDistribution() (icys *model.AdapterIcy, err error) {
	req, err := f.makeReq("/api/v1/projects/icy-distribution/weekly", http.MethodGet, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&icys); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}
	return icys, nil
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

func (f *Fortress) CreateGuildScheduledEvent(e *model.DiscordEvent) error {
	jsonValue, err := json.Marshal(e)
	if err != nil {
		return err
	}

	req, err := f.makeReq("/api/v1/discords/scheduled-events", http.MethodPost, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
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

func (f *Fortress) UpsertRollupRecord(record *model.EngagementsRollupRecord) error {
	jsonValue, err := json.Marshal(record)
	req, err := f.makeReq("/api/v1/engagements/rollup", http.MethodPost, bytes.NewBuffer(jsonValue))
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

func (f *Fortress) CreateBraineryPost(post *model.CreateBraineryLogRequest) error {
	jsonValue, err := json.Marshal(post)
	req, err := f.makeReq("/api/v1/brainery-logs", http.MethodPost, bytes.NewBuffer(jsonValue))
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

func (f *Fortress) GetBraineryReport(view string, date string) (report *model.BraineryMetric, err error) {
	req, err := f.makeReq("/api/v1/brainery-logs/metrics", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("view", view)
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid call, code %v", resp.StatusCode)
	}

	var braineryMetricResp model.BraineryMetricResponse

	if err := json.NewDecoder(resp.Body).Decode(&braineryMetricResp); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return &braineryMetricResp.Data, nil
}

func (f *Fortress) GetEmployees(in EmployeeSearch) ([]model.Employee, error) {
	req, err := f.makeReq("/api/v1/discords", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if in.DiscordID != "" {
		q.Add("discord_id", in.DiscordID)
	}
	if in.Email != "" {
		q.Add("email", in.Email)
	}
	if in.Key != "" {
		q.Add("key", in.Key)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid call, code %v", resp.StatusCode)
	}

	var employeeResp model.FortressEmployeeResponse

	if err := json.NewDecoder(resp.Body).Decode(&employeeResp); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return employeeResp.Data, nil
}

func (f *Fortress) GetEmployeesWithMMAScore() ([]model.EmployeeMMAScore, error) {
	req, err := f.makeReq("/api/v1/discords/mma-scores", http.MethodGet, nil)
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

	var employeeResp model.FortressEmployeeMMAScoreResponse

	if err := json.NewDecoder(resp.Body).Decode(&employeeResp); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return employeeResp.Data, nil
}
func (f *Fortress) GetTrendingRepos(spokenLang string, programLang string, dateRange string) (repos *model.AdapterTrend, err error) {
	// req, err := f.makeReq("/api/v1/notion/issues", http.MethodGet, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	// defer resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("invalid call, code %v", resp.StatusCode)
	// }
	// if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
	// 	return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	// }
	// return issues, nil

	// For now use built-in funcs, not external API call
	// TODO: Replace by API call
	resp, err := crawl(spokenLang, programLang, dateRange)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp, &repos); err != nil {
		return nil, fmt.Errorf("an error occurred when unmarshaling, error %v", err.Error())
	}
	return repos, nil
}

// Parse numerical text "num" to uint16 text, also handle comma-seperated number.
func parseCommaNumber(num string) (uint16, error) {
	num = strings.Replace(num, ",", "", -1)
	parsedNum, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	return uint16(parsedNum), nil
}

// Parse a string in the format '<NUMBER> stars today/last week/last month.' and convert the numerical part into a uint16 value. If the initial portion of the string is not a valid numerical representation, return an error."
func parseStarGained(s string) (uint16, error) {
	s = strings.TrimSpace(s)
	// TODO: handle if empty string
	s = strings.Split(s, " ")[0]
	// Try to convert to uint
	count, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint16(count), nil
}

// Mock API functionality
type MockAPIResponse struct {
	Data []*model.Repo
}

func crawl(spokenLang string, programmingLang string, dateRange string) ([]byte, error) {
	var (
		BASE_URL            = "https://github.com"
		GITHUB_TRENDING_URL = "https://github.com/trending/%s?since=%s&spoken_language_code=%s"
	)
	var repos []*model.Repo
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("github.com", "www.github.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./github_cache"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})
	c.OnHTML("div.Box > div > article.Box-row", func(e *colly.HTMLElement) {
		r := model.Repo{}
		url := BASE_URL + e.ChildAttr("h2 > a", "href")

		starCountText := strings.TrimSpace(e.ChildText("div > span + a"))
		forkCountText := strings.TrimSpace(e.ChildText("div > span + a + a"))
		starGainedText := e.ChildText("div > span + a + a + span + span")
		description := e.ChildText("h2 + p")

		// StarGained: x stars today/last week/last month
		starGained, _ := parseStarGained(starGainedText)
		starCount, _ := parseCommaNumber(starCountText)
		forkCount, _ := parseCommaNumber(forkCountText)

		r.StarGained = starGained
		r.StarCount = starCount
		r.ForkCount = forkCount
		r.ProgrammingLanguage = programmingLang
		r.DateRange = dateRange
		r.Description = description
		r.SpokenLanguage = spokenLang
		r.URL = url
		r.Name = strings.Replace(r.URL, BASE_URL+"/", "", -1)
		repos = append(repos, &r)
	})
	// Start scraping on Github Trending
	c.Visit(
		fmt.Sprintf(
			GITHUB_TRENDING_URL,
			programmingLang, dateRange, spokenLang,
		),
	)
	resp := &MockAPIResponse{}
	// Limit to maximum of 10 repos
	if len(repos) < 10 {
		resp.Data = repos
	} else {
		resp.Data = repos[:10]
	}
	result, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	// Return the scraped repos encoded in json
	return result, nil
}

func (f *Fortress) SalaryAdvance(discordID, amount string) (salaryAdvance *model.AdapterSalaryAdvance, err error) {
	type Request struct {
		DiscordID string `json:"discordID"`
		Amount    string `json:"amount"`
	}

	type ErrResponse struct {
		Data  any    `json:"data"`
		Error string `json:"error"`
	}

	errResponse := ErrResponse{}

	request := Request{
		DiscordID: discordID,
		Amount:    amount,
	}
	jsonValue, _ := json.Marshal(request)

	req, err := f.makeReq("/api/v1/discords/advance-salary", http.MethodPost, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
		}
		return nil, fmt.Errorf(errResponse.Error)
	}

	if err := json.NewDecoder(resp.Body).Decode(&salaryAdvance); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return salaryAdvance, nil
}

func (f *Fortress) CheckAdvanceSalary(discordID string) (salaryAdvance *model.AdapterCheckSalaryAdvance, err error) {
	type Request struct {
		DiscordID string `json:"discordID"`
	}

	type ErrResponse struct {
		Data  any    `json:"data"`
		Error string `json:"error"`
	}

	errResponse := ErrResponse{}

	request := Request{
		DiscordID: discordID,
	}
	jsonValue, _ := json.Marshal(request)

	req, err := f.makeReq("/api/v1/discords/check-advance-salary", http.MethodPost, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
		}
		return nil, fmt.Errorf(errResponse.Error)
	}

	if err := json.NewDecoder(resp.Body).Decode(&salaryAdvance); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return salaryAdvance, nil
}

func (f *Fortress) SalaryAdvanceReport() (unpaidSalaryAdvances *model.AdapterSalaryAdvanceReport, err error) {
	req, err := f.makeReq("/api/v1/discords/salary-advance-report?isPaid=false", http.MethodGet, nil)
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

	if err := json.NewDecoder(resp.Body).Decode(&unpaidSalaryAdvances); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return unpaidSalaryAdvances, nil
}

func (f *Fortress) GetIcyAccounting() (icyAccounting *model.AdapterIcyAccounting, err error) {
	req, err := f.makeReq("/api/v1/discords/icy-accounting", http.MethodGet, nil)
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

	if err := json.NewDecoder(resp.Body).Decode(&icyAccounting); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return icyAccounting, nil
}

func (f *Fortress) ListICYEarnedTransactions(discordID string, page, size int) (*model.AdapterICYEarnedTransactions, error) {
	req, err := f.makeReq(fmt.Sprintf("/api/v1/discords/%s/earns/transactions?page=%d&size=%d", discordID, page, size), http.MethodGet, nil)
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

	var transactions model.AdapterICYEarnedTransactions
	if err := json.NewDecoder(resp.Body).Decode(&transactions); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return &transactions, nil
}

func (f *Fortress) GetICYTotalEarned(discordID string) (*model.AdapterICYTotalEarned, error) {
	req, err := f.makeReq(fmt.Sprintf("/api/v1/discords/%s/earns/total", discordID), http.MethodGet, nil)
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

	var total model.AdapterICYTotalEarned
	if err := json.NewDecoder(resp.Body).Decode(&total); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return &total, nil
}

func (f *Fortress) Get30daysTotalReward() (*model.AdapterICYTotalEarned, error) {
	req, err := f.makeReq("/api/v1/discords/earns/total", http.MethodGet, nil)
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

	var total model.AdapterICYTotalEarned
	if err := json.NewDecoder(resp.Body).Decode(&total); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return &total, nil
}

func (f *Fortress) GetBanks(id, bin, swiftCode string) (banks *model.AdapterBank, err error) {
	req, err := f.makeReq("/api/v1/metadata/banks", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id", id)
	q.Add("bin", bin)
	q.Add("swiftCode", swiftCode)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid call, code %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&banks); err != nil {
		return nil, fmt.Errorf("invalid decoded, error %v", err.Error())
	}

	return banks, nil
}
