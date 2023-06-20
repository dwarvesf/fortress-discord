package brainery

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Brainery struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) Service {
	return &Brainery{
		adapter: adapter,
		l:       l,
	}
}

func (e *Brainery) Post(in *model.Brainery) (*model.Brainery, error) {
	res, err := http.Get(in.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	title := "Untitled"

	titles := strings.Split(doc.Find("title").Text(), "-")
	if len(titles) > 0 {
		title = titles[0]
	}

	description, err := e.adapter.OpenAI().SummarizeBraineryPost(in.URL)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	return &model.Brainery{
		Title:       title,
		URL:         in.URL,
		Author:      in.Author,
		Description: description,
		Reward:      in.Reward,
		PublishDate: in.PublishDate,
		Tags:        in.Tags,
		Github:      in.Github,
		DiscordID:   in.DiscordID,
	}, nil
}
