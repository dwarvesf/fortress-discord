package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
)

type Memo struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) MemoCommander {
	return &Memo{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

func (e *Memo) List(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Memo().GetMemos()
	if err != nil {
		e.L.Error(err, "can'e get list of Memo")
		return err
	}

	// 2. render
	return e.view.Memo().List(message, data)
}

func (e *Memo) ListMemoLogs(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Memo().GetMemoLogs()
	if err != nil {
		e.L.Error(err, "can't get list of Memo")
		return err
	}

	// 2. render
	return e.view.Memo().ListMemoLogs(message, data)
}

func (e *Memo) Sync(message *model.DiscordMessage) error {
	targetChannelID := constant.DiscordReadingChannel
	if e.cfg.Env == "dev" {
		targetChannelID = constant.DiscordPlayGroundReadingChannel
	}
	rawFormattedContent := stringutils.FormatString(message.RawContent)

	extractReward := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternIcyReward)
	extractChannelID := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternDiscordChannelID)
	if len(extractChannelID) > 1 {
		return e.view.Error().Raise(message, "There is more than one target channel in your message.")
	}

	if len(extractChannelID) == 1 {
		targetChannelID = extractChannelID[0]
	}

	reward := constant.DefaultMemoReward
	if len(extractReward) > 0 {
		reward = extractReward[0]
	}

	// 1. sync
	data, err := e.svc.Memo().SyncMemoLogs()
	if err != nil {
		return e.view.Error().Raise(message, "Could not sync memo logs.")
	}
	//data := []model.MemoLog{
	//	{
	//		ID:    "36f2f8ca-9722-4267-8337-c2b8edf1fc54",
	//		Title: "Devbox #1: The world before Docker",
	//		URL:   "https://memo.d.foundation/playground/_memo/devbox-a-world-before-docker/",
	//		Authors: []model.MemoLogAuthor{
	//			{
	//				EmployeeID: "",
	//				GithubID:   "",
	//				DiscordID:  "686038111217909809",
	//			},
	//			{
	//				EmployeeID: "",
	//				GithubID:   "",
	//				DiscordID:  "797042642600722473",
	//			},
	//		},
	//		Description: "test description",
	//		PublishedAt: &time.Time{},
	//		Reward:      decimal.New(10, 0),
	//	},
	//	{
	//		ID:          "43c82b06-b7dc-48c4-90c0-135d211b22aa",
	//		Title:       "Design less, present more with Deckset",
	//		URL:         "https://memo.d.foundation/playground/_memo/design-less-present-more-with-deckset./",
	//		Authors:     []model.MemoLogAuthor{},
	//		Description: "In this March, we're eyeing on what's brewing in the tech market, ICY updates in 2024, the first offline meetup and product demo.",
	//		PublishedAt: &time.Time{},
	//		Reward:      decimal.New(10, 0),
	//	},
	//}

	// 2. render
	return e.view.Memo().Sync(message, data, targetChannelID, reward)
}
