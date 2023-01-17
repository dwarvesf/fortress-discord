package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// onMessageCreate is an entry point for discord message create (chat) event
func (d *Discord) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// this will ignore message send from bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// ignore if match some validation
	if !d.shouldParseMessage(m) {
		d.L.Field("message", m.Message.Content).Debug("ignored message")
		return
	}

	// add "is typing" indicator
	d.Session.ChannelTyping(m.ChannelID)

	// pipetrhough message to command
	err := d.Command.Execute(d.parseMessage(m))
	if err != nil {
		d.L.Error(err, "failed to execute command")
	}
}

func (d *Discord) shouldParseMessage(m *discordgo.MessageCreate) bool {
	// ignore all message from bot
	if m.Author.Bot {
		return false
	}

	// ignore message without prefix
	if !strings.HasPrefix(m.Content, d.Cfg.Discord.Prefix) {
		return false
	}

	// ignore non command format
	if len(strings.Split(m.Content, " ")) < 1 {
		return false
	}

	return true
}

func (d *Discord) parseMessage(m *discordgo.MessageCreate) *model.DiscordMessage {
	parsedContent := strings.Split(m.Message.Content, " ")

	// we remove the prefix from parsedContent[0]
	parsedContent[0] = strings.TrimPrefix(parsedContent[0], d.Cfg.Discord.Prefix)

	return &model.DiscordMessage{
		RawContent:  m.Message.Content,
		ContentArgs: parsedContent,
		ChannelId:   m.ChannelID,
		GuildId:     m.GuildID,
		Author:      m.Author,
	}
}
