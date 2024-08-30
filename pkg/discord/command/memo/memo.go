package memo

import (
	"errors"
	"strings"
	"time"

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
	now := time.Now()
	var (
		from       *time.Time
		timeAmount int
		timeUnit   string
	)

	if len(message.ContentArgs) < 3 {
		// Default to the last 7 days
		tempFrom := now.AddDate(0, 0, -7)
		from = &tempFrom
		timeAmount = 7
		timeUnit = "days"
	} else if len(message.ContentArgs) >= 3 {
		durationStr := strings.Join(message.ContentArgs[2:], " ")

		var (
			tempFrom *time.Time
			err      error
		)
		tempFrom, timeAmount, timeUnit, err = stringutils.ParseTimePeriod(now, durationStr)
		if err != nil {
			return err
		}
		from = tempFrom
	} else {
		return errors.New("invalid command format")
	}

	// 1. get data from service
	data, err := e.svc.Memo().GetMemoLogs(from, &now)
	if err != nil {
		e.L.Error(err, "can't get list of Memo")
		return err
	}

	// 2. render
	return e.view.Memo().ListMemoLogs(message, data, timeAmount, timeUnit)
}

func (e *Memo) ListMemoOpenPullRequest(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := e.svc.Memo().GetMemoOpenPullRequest()
	if err != nil {
		e.L.Error(err, "can't get list of Memo Open Pull Request")
		return err
	}

	// 2. render
	return e.view.Memo().ListMemoOpenPullRequest(message, *data)
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

	// 2. render
	return e.view.Memo().Sync(message, data, targetChannelID, reward)
}

func (e *Memo) ListByDiscordID(message *model.DiscordMessage) error {
	discordID := strings.TrimPrefix(message.ContentArgs[1], "<@")
	discordID = strings.TrimSuffix(discordID, ">")

	// Get memo data for the specified Discord ID
	data, err := e.svc.Memo().GetMemosByDiscordID(discordID)
	if err != nil {
		e.L.Error(err, "can't get memos for Discord ID")
		return e.view.Error().Raise(message, "Could not fetch memos for the specified Discord ID.")
	}

	// Render the result
	return e.view.Memo().ListByDiscordID(message, data, discordID)
}
