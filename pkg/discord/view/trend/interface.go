package trend

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TrendViewer interface {
	List(message *model.DiscordMessage, repos []*model.Repo) error
	Help(message *model.DiscordMessage) error
	ListProgramLang(message *model.DiscordMessage) error
	ListSpokenLang(message *model.DiscordMessage) error
	ListDateRange(message *model.DiscordMessage) error
	GetAvailableProgrammingLang() []string
	GetAvailableSpokenLangMap() map[string]string
	GetAvaiableDateRangeMap() map[string]string
	NotFound(message *model.DiscordMessage) error
}
