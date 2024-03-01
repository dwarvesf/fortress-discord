package brainery

import (
	"strings"
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/brainery"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
)

func (e *Brainery) Post(message *model.DiscordMessage) error {
	targetChannelID := constant.DiscordBraineryChannel
	if e.cfg.Env == "dev" {
		targetChannelID = constant.DiscordPlayGroundBraineryChannel
	}
	rawFormattedContent := stringutils.FormatString(message.RawContent)
	now := time.Now()

	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	publishedAt := now.In(loc)

	extractURL := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternUrl)
	extractDiscordID := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternDiscordID)
	extractReward := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternIcyReward)
	extractGithub := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternGithub)
	extractDesc := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternDescription)

	if len(extractURL) == 0 || len(extractURL) > 1 {
		return e.view.Error().Raise(message, "There is no URL or more than one URL in your message.")
	}

	if !strings.Contains(extractURL[0], "https://brain.d.foundation") {
		return e.view.Error().Raise(message, "The article should be get https://brain.d.foundation.")
	}

	if len(extractDiscordID) == 0 || len(extractDiscordID) > 1 {
		return e.view.Error().Raise(message, "There is no valid user or more than one user tagged in your message.")
	}

	extractChannelID := stringutils.ExtractPattern(rawFormattedContent, constant.RegexPatternDiscordChannelID)
	if len(extractChannelID) > 1 {
		return e.view.Error().Raise(message, "There is more than one target channel in your message.")
	}

	if len(extractChannelID) == 1 {
		targetChannelID = extractChannelID[0]
	}

	reward := constant.DefaultBraineryReward
	if len(extractReward) > 0 {
		reward = extractReward[0]
	}

	gh := ""
	if len(extractGithub) > 0 {
		gh = extractGithub[0]
	}

	desc := ""
	if len(extractDesc) > 0 {
		desc = extractDesc[0]
	}

	mbrainery := &brainery.PostInput{
		URL:         extractURL[0],
		DiscordID:   extractDiscordID[0],
		Description: desc,
		Reward:      reward,
		PublishedAt: &publishedAt,
		//Tags:        extractTags,
		Github: gh,
	}

	braineryData, err := e.svc.Brainery().Post(mbrainery)
	if err != nil {
		return e.view.Error().Raise(message, err.Error())
	}
	err = e.view.Brainery().Post(message, braineryData, targetChannelID)
	if err != nil {
		return e.view.Error().Raise(message, err.Error())
	}
	// 2. render
	return nil
}
