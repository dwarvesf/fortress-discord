package changelog

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Changelog struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) ChangelogViewer {
	return &Changelog{
		ses: ses,
	}
}

func (c *Changelog) Changelog(message *model.DiscordMessage, data []*model.Changelog) error {
	content := []string{}
	for i, v := range data {
		text := "**" + strconv.Itoa(i) + " " + v.Name + "** - " + "[" + v.Title + "]" + "(" + v.ChangelogURL + ")"
		content = append(content, text)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Project Changelogs**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(c.ses, message, msg)
}

func (c *Changelog) ChangelogSendSuccess(message *model.DiscordMessage, data *model.Changelog) error {
	content := "Send changelog for " + "**" + data.Title + "** " + "successfully"

	msg := &discordgo.MessageEmbed{
		Title:       "**Project Changelogs**",
		Description: content,
	}

	return base.SendEmbededMessage(c.ses, message, msg)
}

func (c *Changelog) ChangelogSendFailed(message *model.DiscordMessage, data *model.Changelog) error {
	content := "Send changelog for " + "**" + data.Title + "** " + "failed"

	msg := &discordgo.MessageEmbed{
		Title:       "**Project Changelogs**",
		Description: content,
	}

	return base.SendEmbededMessage(c.ses, message, msg)
}
