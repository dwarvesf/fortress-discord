package brainery

import (
	"regexp"
	"strings"
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service/brainery"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Brainery struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) Commander {
	return &Brainery{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

var (
	discordIDRegexPattern   = `<@(\d+)>`
	tagRegexPattern         = `#(\w+)`
	icyRewardRegexPattern   = ` (\d+)`
	urlRegexPattern         = `((?:https?://)[^\s]+)`
	githubRegexPattern      = `gh:(\w+)`
	descriptionRegexPattern = `d:"(.*?)"`
)

const defaultBraineryReward = "0"

func (e *Brainery) Post(message *model.DiscordMessage) error {
	targetChannelID := constant.DiscordBraineryChannel
	if e.cfg.Env == "local" {
		targetChannelID = constant.DiscordPlayGroundBraineryChannel
	}
	rawFormattedContent := formatString(message.RawContent)
	now := time.Now()

	extractURL := extractPattern(rawFormattedContent, urlRegexPattern)
	extractDiscordID := extractPattern(rawFormattedContent, discordIDRegexPattern)
	extractTags := extractPattern(rawFormattedContent, tagRegexPattern)
	extractReward := extractPattern(rawFormattedContent, icyRewardRegexPattern)
	extractGithub := extractPattern(rawFormattedContent, githubRegexPattern)
	extractDesc := extractPattern(rawFormattedContent, descriptionRegexPattern)

	if len(extractURL) == 0 || len(extractURL) > 1 {
		return e.view.Error().Raise(message, "There is no URL or more than one URL in your message.")
	}

	if !strings.Contains(extractURL[0], "https://brain.d.foundation") {
		return e.view.Error().Raise(message, "The article should be get https://brain.d.foundation.")
	}

	if len(extractDiscordID) == 0 || len(extractDiscordID) > 1 {
		return e.view.Error().Raise(message, "There is no valid user or more than one user tagged in your message.")
	}

	reward := defaultBraineryReward
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
		PublishedAt: &now,
		Tags:        extractTags,
		Github:      gh,
	}

	result, err := e.svc.Brainery().Post(mbrainery)
	if err != nil {
		return e.view.Error().Raise(message, err.Error())
	}

	// 2. render
	return e.view.Brainery().Post(message, result, targetChannelID)

}

func extractPattern(str string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(str, -1)

	var result []string
	for _, match := range matches {
		result = append(result, match[1])
	}

	return result
}

func formatString(str string) string {
	// Replace spaces with a single space
	re := regexp.MustCompile(`\s+`)
	formattedStr := re.ReplaceAllString(str, " ")

	// Remove spaces after the "#" symbol
	formattedStr = strings.ReplaceAll(formattedStr, "# ", "#")

	return formattedStr
}
