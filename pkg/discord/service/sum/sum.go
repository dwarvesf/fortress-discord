package sum

import (
	"fmt"
	"strings"

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
	data, err := e.adapter.Dify().SummarizeArticle(url)
	if err != nil {
		fmt.Printf("failed to summarize the given article. Error: %v\n", err)
		return nil, err
	}
	title, summary := extractTitleAndSummary(data)

	return &model.Sum{
		URL:     url,
		Title:   title,
		Summary: summary,
	}, nil
}

func extractTitleAndSummary(input string) (string, string) {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "", ""
	}

	title := lines[0]

	// Filter out empty lines
	var contentLines []string
	for _, line := range lines[1:] {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			contentLines = append(contentLines, trimmedLine)
		}
	}

	summary := strings.Join(contentLines, "\n")
	return title, summary
}
