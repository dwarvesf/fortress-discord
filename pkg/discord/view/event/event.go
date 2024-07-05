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
		`**?event list/ls**・get discord scheduled events
		**?event speakerset spks <eventID> <topic1> @user1 @user2 <topic2> @user3**・set speakers for scheduled events
		`,
		"*Example:* `?event spks 1256184313876054147 Erlang @bienvh AI-101 @cherry`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, message, msg)
}
