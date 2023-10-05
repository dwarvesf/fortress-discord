package brainery

import (
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
	"strings"
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/brainery"
	"github.com/dwarvesf/fortress-discord/pkg/model"
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

	extractURL := stringutils.ExtractPattern(rawFormattedContent, constant.UrlRegexPattern)
	extractDiscordID := stringutils.ExtractPattern(rawFormattedContent, constant.DiscordIDRegexPattern)
	// TODO: need to change regex pattern to detect tag without the conflict with channel pattern
	//extractTags := extractPattern(rawFormattedContent, tagRegexPattern)
	extractReward := stringutils.ExtractPattern(rawFormattedContent, constant.IcyRewardRegexPattern)
	extractGithub := stringutils.ExtractPattern(rawFormattedContent, constant.GithubRegexPattern)
	extractDesc := stringutils.ExtractPattern(rawFormattedContent, constant.DescriptionRegexPattern)

	if len(extractURL) == 0 || len(extractURL) > 1 {
		return e.view.Error().Raise(message, "There is no URL or more than one URL in your message.")
	}

	if !strings.Contains(extractURL[0], "https://brain.d.foundation") {
		return e.view.Error().Raise(message, "The article should be get https://brain.d.foundation.")
	}

	if len(extractDiscordID) == 0 || len(extractDiscordID) > 1 {
		return e.view.Error().Raise(message, "There is no valid user or more than one user tagged in your message.")
	}

	extractChannelID := stringutils.ExtractPattern(rawFormattedContent, constant.DiscordChannelIDRegexPattern)
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
