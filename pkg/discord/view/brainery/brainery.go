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
	latestPost := "This is where we keep track of our **top 10** latest Brainery notes:\n\n"
	if braineryMetric.LatestPosts != nil {
		for _, post := range braineryMetric.LatestPosts {
			latestPost += fmt.Sprintf("•  [%s](%s)\n", post.Title, post.URL)
		}
	}

	contributors := ""
	if len(braineryMetric.Contributors) > 0 {
		for _, itm := range braineryMetric.Contributors {
			totalICY = totalICY.Add(itm.Reward)
			contributors += fmt.Sprintf("• <@%v> - [%s](%s)\n", itm.DiscordID, itm.Title, itm.URL)
		}
	}

	if len(contributors) > 0 {
		latestPost += "**\nContributors**\n"
		latestPost += contributors
	}

	newContributor := ""
	if len(braineryMetric.NewContributors) > 0 {
		for _, itm := range braineryMetric.NewContributors {
			totalICY = totalICY.Add(itm.Reward)
			newContributor += fmt.Sprintf("• <@%v> - [%s](%s)\n", itm.DiscordID, itm.Title, itm.URL)
		}
	}

	if len(newContributor) > 0 {
		latestPost += "\n**New Contributors**\n"
		latestPost += newContributor
	}

	if totalICY.GreaterThan(decimal.NewFromInt(0)) {
		embedField := &discordgo.MessageEmbedField{
			Name:   "Total ICY Given Out",
			Value:  totalICY.String() + " ICY 🧊",
			Inline: false,
		}

		messageEmbed = append(messageEmbed, embedField)
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
		Title:       fmt.Sprintf("BRAINERY %s REPORT ", strings.ToTitle(view)),
		Fields:      messageEmbed,
		Description: latestPost,
	}

	return base.SendEmbededMessageWithChannel(v.ses, original, msg, channelID)
}
