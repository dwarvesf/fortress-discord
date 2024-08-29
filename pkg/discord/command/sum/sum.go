package sum

import (
	"fmt"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Sum struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) SumCommander {
	return &Sum{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Sum) Sum(message *model.DiscordMessage) error {
	var template, url string
	var data *model.Sum
	var err error

	switch len(message.ContentArgs) {
	case 2:
		// Original behavior: ?sum <url>
		url = message.ContentArgs[1]
		data, err = e.svc.Sum().SummarizeArticle("", url)
	case 3:
		// New behavior: ?sum <template> <url>
		template = message.ContentArgs[1]
		url = message.ContentArgs[2]
		data, err = e.svc.Sum().SummarizeArticle(template, url)
	default:
		errorSummary := &model.Sum{
			Title:   "Error",
			Summary: "Invalid command format. Use: ?sum <url> or ?sum <template> <url>",
		}
		return e.view.Sum().Sum(message, errorSummary)
	}

	if err != nil {
		e.L.Error(err, "failed to summarize the given article")
		errorSummary := &model.Sum{
			Title:   "Error",
			Summary: fmt.Sprintf("Failed to summarize the article: %v", err),
		}
		return e.view.Sum().Sum(message, errorSummary)
	}

	// render
	return e.view.Sum().Sum(message, data)
}
