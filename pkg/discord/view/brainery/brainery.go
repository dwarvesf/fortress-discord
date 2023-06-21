package brainery

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"

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

func (h *Brainery) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?brainery post**・publish new brainery article.",
		"*Example:* `?brainery post <url> @n #tag1 #tag2 gh:namnhce`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(h.ses, message, msg)
}

func (h *Brainery) Post(original *model.DiscordMessage, content *model.Brainery, channelID string) error {
	author, err := h.ses.GuildMember(original.GuildId, content.DiscordID)
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

		Footer: &discordgo.MessageEmbedFooter{
			Text: "Added at " + content.PublishedAt.Format("January 2, 2006 3:04 PM") + " 🎉🎉🎉",
		},
		Timestamp: "custom",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatar,
		},
	}

	return base.SendEmbededMessageWithChannel(h.ses, original, msg, channelID)
}
