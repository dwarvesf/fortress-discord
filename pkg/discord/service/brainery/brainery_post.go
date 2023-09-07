package brainery

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/shopspring/decimal"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Brainery) Post(in *PostInput) (*model.Brainery, error) {
	title, err := getTitle(in.URL)
	if err != nil {
		return nil, err
	}

	githubRawURL, err := convertURL(in.URL)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(githubRawURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	metadata := extractMetadata(string(body)[:300])
	tags := buildTags(in.Tags)
	if tags == "" {
		tags = buildTags(metadata.Tags)
		if tags == "" {
			return nil, fmt.Errorf("There is no tags in metadata.\nInput tags manually like this format #tag1 #tag2")
		}
	}

	if in.Github == "" {
		in.Github = metadata.GithubID
	}

	if in.Reward == "0" {
		in.Reward = "10"
		if metadata.Icy != "" {
			in.Reward = metadata.Icy
		}

	}

	description := in.Description
	if description == "" {
		maxlength := 3000
		if len(string(body)) < maxlength {
			maxlength = len(string(body))
		}

		description, err = e.adapter.OpenAI().SummarizeBraineryPost(string(body)[:maxlength])
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return nil, err
		}
	}

	mBrainery := &model.Brainery{
		Title:       title,
		URL:         in.URL,
		Author:      in.Author,
		Description: description + "..",
		Reward:      in.Reward,
		PublishedAt: in.PublishedAt,
		Tags:        tags,
		Github:      in.Github,
		DiscordID:   in.DiscordID,
	}

	rw, err := decimal.NewFromString(mBrainery.Reward)
	if err != nil {
		return nil, err
	}

	err = e.adapter.Fortress().CreateBraineryPost(&model.CreateBraineryLogRequest{
		Title:       mBrainery.Title,
		URL:         mBrainery.URL,
		GithubID:    mBrainery.Github,
		DiscordID:   mBrainery.DiscordID,
		Tags:        separateTags(mBrainery.Tags),
		PublishedAt: mBrainery.PublishedAt.Format(time.RFC3339),
		Reward:      rw,
	})
	if err != nil {
		return nil, err
	}

	return mBrainery, nil
}

func convertURL(originalURL string) (string, error) {
	filepath := strings.TrimPrefix(originalURL, "https://brain.d.foundation/")
	filepath = strings.ReplaceAll(filepath, "+", " ")

	// Build the new URL with the desired format
	newURL := fmt.Sprintf("https://raw.githubusercontent.com/dwarvesf/brain/master/%s.md", filepath)
	parsedURL, err := url.Parse(newURL)
	if err != nil {
		return "", err
	}
	return parsedURL.String(), nil
}

func getTitle(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(doc.Find("title").Text(), " - The Dwarves Brainery"), nil
}

type Metadata struct {
	Tags     []string `yaml:"tags"`
	Author   string   `yaml:"author"`
	GithubID string   `yaml:"github_id"`
	Date     string   `yaml:"date"`
	Icy      string   `yaml:"icy"`
}

func extractMetadata(content string) *Metadata {
	// Find the start and end positions of the metadata section
	start := strings.Index(content, "---")
	end := strings.Index(content[start+3:], "---")

	var data Metadata

	if start != -1 && end != -1 {
		metadataSection := content[start+3 : start+3+end]

		err := yaml.Unmarshal([]byte(metadataSection), &data)
		if err != nil {
			return nil
		}
	}

	return &data
}

func buildTags(tags []string) string {
	var result []string
	for _, tag := range tags {
		if tag == "" {
			continue
		}
		result = append(result, "#"+tag)
	}

	return strings.Join(result, "\n")
}
func separateTags(tags string) []string {
	return strings.Split(strings.ReplaceAll(tags, "#", ""), "\n")
}
