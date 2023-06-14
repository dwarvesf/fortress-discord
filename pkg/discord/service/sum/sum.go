package sum

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Sum struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) SumServicer {
	return &Sum{
		adapter: adapter,
		l:       l,
	}
}

func (e *Sum) SummarizeArticle(url string) (*model.Sum, error) {
	res, err := http.Get(url)

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

	title := doc.Find("title").Text()

	data, err := e.adapter.OpenAI().SummarizeArticle(url)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	return &model.Sum{
		URL:     url,
		Title:   title,
		Summary: data,
	}, nil
}
