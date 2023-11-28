package trend

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

const DEFAULT_SPOKEN_LANGUAGE = "en"
const DEFAULT_DATE_RANGE = "daily"
const DEFAULT_PROGRAMMING_LANGUAGE = "go"

type Trend struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) TrendCommander {
	return &Trend{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (e *Trend) DefaultTrend(message *model.DiscordMessage) error {
	data, err := e.svc.Trend().GetTrendingRepos("en", "go", "daily")
	if err != nil {
		e.L.Error(err, "can't get github trending repos")
		return err
	}
	// 2. render
	return e.view.Trend().List(message, data)
}
func (e *Trend) Trend(message *model.DiscordMessage) error {
	// 1. get data from service
	var spokenLang, programmingLang, dateRange string
	switch len(message.ContentArgs) {
	case 2:
		switch message.ContentArgs[1] {
		case "programming":
			return e.view.Trend().ListProgramLang(message)
		case "date":
			return e.view.Trend().ListDateRange(message)
		case "spoken":
			return e.view.Trend().ListSpokenLang(message)
		default:
			spokenLang = DEFAULT_SPOKEN_LANGUAGE
			dateRange = DEFAULT_DATE_RANGE
			programmingLang = message.ContentArgs[1]
		}
	case 3:
		spokenLang = DEFAULT_SPOKEN_LANGUAGE
		programmingLang = message.ContentArgs[1]
		dateRange = message.ContentArgs[2]
	case 4:
		programmingLang = message.ContentArgs[1]
		dateRange = message.ContentArgs[2]
		spokenLang = message.ContentArgs[3]
	default:
		e.L.Error(nil, "Unexpected params for command. Expect ?trend <spoken_lang> <program_lang> <date_range>")
	}
	data, err := e.svc.Trend().GetTrendingRepos(spokenLang, programmingLang, dateRange)

	if err != nil {
		e.L.Error(err, "can't get github trending repos")
		return err
	}
	if len(data) == 0 {
		return e.view.Trend().NotFound(message)
	}
	// 2. render
	return e.view.Trend().List(message, data)
}
