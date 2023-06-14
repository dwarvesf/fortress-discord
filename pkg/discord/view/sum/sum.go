package sum

import (
	"fmt"

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
		Title:       fmt.Sprintf("%s", summary.Title),
		Description: summary.Summary,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
