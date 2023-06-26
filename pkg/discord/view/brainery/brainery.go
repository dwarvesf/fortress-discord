package brainery

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shopspring/decimal"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Brainery struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &Brainery{
		ses: ses,
	}
}

func (v *Brainery) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?brainery post**・publish new brainery article.",
		"*Example:* `?brainery post <url> @n #tag1 #tag2 gh:namnhce`",
		"**?brainery report**・get brainery report by week/month.",
		"*Example:* `?brainery report weekly`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}

func (v *Brainery) Post(original *model.DiscordMessage, content *model.Brainery, channelID string) error {
	author, err := v.ses.GuildMember(original.GuildId, content.DiscordID)
	if err != nil {
		return err
	}

	avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.webp?size=240", author.User.ID, author.User.Avatar)
	if author.Avatar != "" {
		avatar = fmt.Sprintf("https://cdn.discordapp.com/guilds/%v/users/%v/avatars/%v.webp?size=240", original.GuildId, author.User.ID, author.Avatar)
	}

	authorField := fmt.Sprintf("<@%s>", author.User.ID)
	if content.Github != "" {
		authorField = fmt.Sprintf("<@%s> \n\n**Github**\n[%v](https://github.com/%v)", author.User.ID, content.Github, content.Github)
	}

	messageEmbed := []*discordgo.MessageEmbedField{
		{
			Name:   "Author",
			Value:  authorField,
			Inline: true,
		},
		{
			Name:   "Tags",
			Value:  content.Tags,
			Inline: true,
		},
		{
			Name:   "ICY 🧊",
			Value:  content.Reward,
			Inline: true,
		},
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s", content.Title),
		URL:         content.URL,
		Description: content.Description,
		Fields:      messageEmbed,
		Timestamp:   content.PublishedAt.Format(time.RFC3339),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatar,
		},
	}

	return base.SendEmbededMessageWithChannel(v.ses, original, msg, channelID)
}

func (v *Brainery) Report(original *model.DiscordMessage, view string, braineryMetric *model.BraineryMetric, channelID string) error {
	var messageEmbed []*discordgo.MessageEmbedField
	totalICY := decimal.NewFromInt(0)
	content := ""

	var newBraineryPost []model.BraineryMetricItem
	newBraineryPost = append(newBraineryPost, braineryMetric.Contributors...)
	newBraineryPost = append(newBraineryPost, braineryMetric.NewContributors...)

	if len(newBraineryPost) == 0 {
		content += fmt.Sprintf("There is no new brainery note in this period. This is where we keep track of our **top 10 latest** Brainery notes:\n\n")

		for _, itm := range braineryMetric.LatestPosts {
			content += fmt.Sprintf("• [%s](%s) <@%v>\n", itm.Title, itm.URL, itm.DiscordID)
		}

	} else {
		newBraineryPostStr := ""
		for _, itm := range newBraineryPost {
			totalICY = totalICY.Add(itm.Reward)
			newBraineryPostStr += fmt.Sprintf("• [%s](%s) <@%v>\n", itm.Title, itm.URL, itm.DiscordID)
		}

		if len(newBraineryPostStr) > 0 {
			content += "**Latest Notes** :fire::fire::fire:\n"
			content += newBraineryPostStr + "\n"
		}
	}

	if view == "monthly" {
		topContributor := calculateTopContributor(braineryMetric.TopContributors)
		content += topContributor + "\n"
	}

	newContributor := ""
	if len(braineryMetric.NewContributors) > 0 {
		ids := make(map[string]bool)
		for _, itm := range braineryMetric.NewContributors {
			v, ok := ids[itm.DiscordID]
			if ok && v {
				continue
			}
			ids[itm.DiscordID] = true
			newContributor += fmt.Sprintf("<@%v> ", itm.DiscordID)
		}
	}

	if newContributor != "" {
		content += fmt.Sprintf("**New Contributors**\n")
		content += newContributor + "\n"
	}

	if totalICY.GreaterThan(decimal.NewFromInt(0)) {
		content += fmt.Sprintf("\n**Total Reward Distributed**\n")
		content += totalICY.String() + " ICY 🧊"
	}

	tags := ""
	if len(braineryMetric.Tags) > 0 {
		for _, tag := range braineryMetric.Tags {
			tags += fmt.Sprintf("#%v ", tag)
		}
	}

	if len(tags) > 0 {
		embedField := &discordgo.MessageEmbedField{
			Name:   "Tags",
			Value:  tags,
			Inline: false,
		}

		messageEmbed = append(messageEmbed, embedField)
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("BRAINERY %s REPORT", strings.ToTitle(view)),
		Fields:      messageEmbed,
		Description: content,
	}

	return base.SendEmbededMessageWithChannel(v.ses, original, msg, channelID)
}

func calculateTopContributor(topContributors []model.TopContributor) string {
	topContributorStr := ""
	if len(topContributors) == 0 {
		return ""
	}

	countMap := make(map[int][]string)
	var uniqueCounts []int

	for _, contributor := range topContributors {
		if contributor.Count > 1 {
			count := contributor.Count
			discordID := contributor.DiscordID
			countMap[count] = append(countMap[count], discordID)

			// Check if count is already in uniqueCounts
			found := false
			for _, uniqueCount := range uniqueCounts {
				if uniqueCount == count {
					found = true
					break
				}
			}

			// If count is not found, add it to uniqueCounts
			if !found {
				uniqueCounts = append(uniqueCounts, count)
			}
		}
	}

	emojiMap := map[int]string{
		0: ":first_place:",
		1: ":second_place:",
		2: ":third_place:",
	}

	// Iterate over uniqueCounts to access Discord IDs in order
	for idx, count := range uniqueCounts {
		discordIDs := countMap[count]
		discordIDStr := ""
		for i := 0; i < len(discordIDs); i++ {
			discordIDStr += "<@" + discordIDs[i] + ">, "
		}

		emojiIdx := idx
		if idx > 2 {
			emojiIdx = 2
		}

		topContributorStr += fmt.Sprintf("%v %v (x%v) \n", emojiMap[emojiIdx], strings.TrimSuffix(discordIDStr, ", "), count)
	}

	topContributor := ""
	if len(topContributorStr) > 0 {
		topContributor += "**Top Contributors**\n"
		topContributor += topContributorStr
	}

	return topContributor
}
