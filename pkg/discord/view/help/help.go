package help

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Help struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) HelpViewer {
	return &Help{
		ses: ses,
	}
}

func (h *Help) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?earn**・earn $ICY for free",
		"**?trial**・list of trial tech",
		"**?assess**・list of assess tech",
		"**?adopt**・list of adopt tech",
		"**?hold**・list of on-hold tech",
		"**?new**・list of new subscribers",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(h.ses, message, msg)
}
