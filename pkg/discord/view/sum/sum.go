package sum

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Sum struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) SumViewer {
	return &Sum{
		ses: ses,
	}
}

func (e *Sum) Sum(original *model.DiscordMessage, summary *model.Sum) error {
	msg := &discordgo.MessageEmbed{
		Title:       summary.Title,
		Description: summary.Summary,
		URL:         summary.URL,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
