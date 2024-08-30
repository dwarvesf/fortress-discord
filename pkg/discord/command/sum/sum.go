package sum

import (
	"fmt"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"net/url"
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
	var template string
	var articleURL string
	var data *model.Sum
	var err error

	// Parse arguments
	for _, arg := range message.ContentArgs[1:] {
		if isValidURL(arg) {
			articleURL = arg
		} else if model.IsValidTemplateType(arg) {
			template = arg
		}
	}

	// Validate parsed arguments
	if articleURL == "" {
		errorSummary := &model.Sum{
			Title:   "Error",
			Summary: "No valid URL provided. Use: ?sum <url> or ?sum <template> <url> or ?sum <url> <template>",
		}
		return e.view.Sum().Sum(message, errorSummary)
	}

	// If no template is provided, use the default summary template
	if template == "" {
		template = string(model.TemplateSummary)
	}

	// Summarize the article
	data, err = e.svc.Sum().SummarizeArticle(template, articleURL)
	if err != nil {
		e.L.Error(err, "failed to summarize the given article")
		errorSummary := &model.Sum{
			Title:   "Error",
			Summary: fmt.Sprintf("Failed to summarize the article: %v", err),
		}
		return e.view.Sum().Sum(message, errorSummary)
	}

	// Render the summary
	return e.view.Sum().Sum(message, data)
}

// isValidURL checks if a string is a valid URL
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
