package event

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Event struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) EventViewer {
	return &Event{
		ses: ses,
	}
}

func (e *Event) Help(message *model.DiscordMessage) error {
	content := []string{
		`**?event list**・get events from notion
		**?event scheduled list/ls**・get discord scheduled events
		**?event scheduled set/s speaker/spk <discord_event_id> @user1:topic1 @user2:topic2**・set speakers for scheduled events
		`,
		"*Example:* `?event scheduled s spk 123871623 @nam:topic1 @nam2:topic2 @nam3:topic3`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
