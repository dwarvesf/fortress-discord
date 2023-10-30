package trend

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func findEmoji(session *discordgo.Session, emojiString string, guildId string, isAnimated bool) string {
	allEmojis, _ := session.GuildEmojis(guildId)
	for _, emoji := range allEmojis {
		if emoji.Name == emojiString {
			if isAnimated {
				return fmt.Sprintf("<a:%s:%s>", emojiString, emoji.ID)
			} else {
				return fmt.Sprintf("<:%s:%s", emojiString, emoji.ID)
			}
		}
	}
	return ""
}

var numberEmojiStrings = [10]string{"", "", "", "four", "five", "six", "seven", "eight", "nine", "keycap_ten"}
var dateRangeStarGainedMap = map[string]string{"daily": "Day", "weekly": "Week", "monthly": "Month"}
var spokenLangMap = map[string]string{"en": "English", "zh": "Chinese", "ru": "Russian"}
var programmingLanguages = []string{"go", "c", "cpp", "javascript", "java", "elixir", "dart", "rust"}

func (e *Trend) List(message *model.DiscordMessage, repos []*model.Repo) error {
	var (
		badge1       = findEmoji(e.ses, "badge1", message.GuildId, true)
		badge2       = findEmoji(e.ses, "badge2", message.GuildId, true)
		badge3       = findEmoji(e.ses, "badge3", message.GuildId, true)
		starAnimated = findEmoji(e.ses, "star_animated", message.GuildId, true)
	)
	numberEmojiStrings[0] = badge1
	numberEmojiStrings[1] = badge2
	numberEmojiStrings[2] = badge3
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"
	var content string
	if len(repos) != 0 {
		content += fmt.Sprintf("Trending repo in %s, %s, %s \n\n", cases.Title(language.AmericanEnglish).String(repos[0].ProgrammingLanguage), spokenLangMap[repos[0].SpokenLanguage], dateRangeStarGainedMap[repos[0].DateRange])
		// Set star gained text(today/last week/last month)
		for i := range repos {
			repo := repos[i]
			// Top 3 repos will have bigger title
			if i < 3 {
				content += fmt.Sprintf("%s [**%s**](%s)\n", numberEmojiStrings[i], repo.Name, repo.URL)

			} else {
				content += fmt.Sprintf("%s [**%s**](%s)\n", fmt.Sprintf(":%s:", numberEmojiStrings[i]), repo.Name, repo.URL)

			}
			truncatedDescription, _ := Truncate(repo.Description, 60)
			content += fmt.Sprintf("*%s*\n%s`%s: %s  |  Total: %s`\n",
				truncatedDescription,
				starAnimated,
				dateRangeStarGainedMap[repos[0].DateRange],
				rightPadding(fmt.Sprint(repo.StarGained), 3),
				fmt.Sprint(repo.StarCount),
			)
		}
	}
	msg := &discordgo.MessageEmbed{
		Title:       "",
		Description: title + content,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Trend) ListDateRange(message *model.DiscordMessage) error {
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"
	var messageEmbed []*discordgo.MessageEmbedField
	for k := range dateRangeStarGainedMap {
		embedField := &discordgo.MessageEmbedField{
			Name: k,
		}
		messageEmbed = append(messageEmbed, embedField)
	}
	msg := &discordgo.MessageEmbed{
		Description: title + "Available value for <date_range> parameter",

		Fields: messageEmbed,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Trend) ListProgramLang(message *model.DiscordMessage) error {
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"

	var messageEmbed []*discordgo.MessageEmbedField
	for _, v := range programmingLanguages {
		embedField := &discordgo.MessageEmbedField{
			Name: v,
		}
		messageEmbed = append(messageEmbed, embedField)
	}
	msg := &discordgo.MessageEmbed{
		Description: title + "Available value for <lang> parameter",
		Fields:      messageEmbed,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func (e *Trend) ListSpokenLang(message *model.DiscordMessage) error {
	title := "### :chart_with_upwards_trend: Github Trending Repositories\n"

	var messageEmbed []*discordgo.MessageEmbedField
	for k := range spokenLangMap {
		embedField := &discordgo.MessageEmbedField{
			Name: k,
		}
		messageEmbed = append(messageEmbed, embedField)
	}
	msg := &discordgo.MessageEmbed{
		Description: title + "Available value for <spoken_lang> parameter",
		Fields:      messageEmbed,
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}

func rightPadding(value interface{}, n int) string {
	// Convert the value to a string
	valueStr := fmt.Sprint(value)

	// Calculate the padding count
	paddingCount := n - len(valueStr)
	if paddingCount < 0 {
		paddingCount = 0
	}

	// Pad the value and return
	paddedValue := valueStr + strings.Repeat(" ", paddingCount)
	return paddedValue
}
func leftPadding(value interface{}, n int) string {
	// Convert the value to a string
	valueStr := fmt.Sprint(value)

	// Calculate the padding count
	paddingCount := n - len(valueStr)
	if paddingCount < 0 {
		paddingCount = 0
	}

	// Pad the value and return
	paddedValue := strings.Repeat(" ", paddingCount) + valueStr
	return paddedValue
}

func (e *Trend) GetAvailableProgrammingLang() []string {
	return programmingLanguages
}
func (e *Trend) GetAvailableSpokenLangMap() map[string]string {
	return spokenLangMap
}
func (e *Trend) GetAvaiableDateRangeMap() map[string]string {
	return dateRangeStarGainedMap
}
func Truncate(text string, width int) (string, error) {
	if width < 0 {
		return "", fmt.Errorf("invalid width size")
	}
	if width >= len(text) {
		return text, nil
	}
	r := []rune(text)

	trunc := r[:width]
	return string(trunc) + "...", nil
}
