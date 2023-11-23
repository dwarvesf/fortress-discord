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

	// pipethrough message to command
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
	args := strings.Split(m.Content, " ")
	if len(args) < 1 {
		return false
	}

	//
	if len(args[0]) == 1 {
		return false
	}

	if strings.ReplaceAll(args[0], d.Cfg.Discord.Prefix, "") == "" {
		return false
	}

	return true
}

func (d *Discord) parseMessage(m *discordgo.MessageCreate) *model.DiscordMessage {
	parsedContent := strings.Split(m.Message.Content, " ")

	// we remove the prefix from parsedContent[0]
	parsedContent[0] = strings.TrimPrefix(parsedContent[0], d.Cfg.Discord.Prefix)

	// we paste the roles to in-app  struct, make sure nil check
	var roles []string
	if m.Message.Member != nil && m.Message.Member.Roles != nil {
		roles = m.Message.Member.Roles
	}

	return &model.DiscordMessage{
		RawContent:  m.Message.Content,
		ContentArgs: parsedContent,
		ChannelId:   m.ChannelID,
		GuildId:     m.GuildID,
		Author:      m.Author,
		Roles:       roles,
	}
}

func (d *Discord) onReactionCreate(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		l := d.L.AddField("channelID", m.ChannelID)
		l.Error(err, "unable to get channel")
		return
	}
	record := &model.EngagementsRollupRecord{
		DiscordUserID: m.UserID,
		LastMessageID: m.MessageID,
		ChannelID:     channel.ID,
		CategoryID:    channel.ParentID,
		MessageCount:  0,
		ReactionCount: 1,
	}
	l := d.L.AddField("record", record)
	err = d.Command.S.Engagement().UpsertRollup(record)
	if err != nil {
		l.Error(err, "unable to upsert record")
		return
	}
	l.Info("increased reaction count")
}

func (d *Discord) onReactionRemove(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		l := d.L.AddField("channelID", m.ChannelID)
		l.Error(err, "unable to get channel")
		return
	}
	record := &model.EngagementsRollupRecord{
		DiscordUserID: m.UserID,
		LastMessageID: m.MessageID,
		ChannelID:     channel.ID,
		CategoryID:    channel.ParentID,
		MessageCount:  0,
		ReactionCount: -1,
	}
	l := d.L.AddField("record", record)
	err = d.Command.S.Engagement().UpsertRollup(record)
	if err != nil {
		l.Error(err, "unable to upsert record")
		return
	}
	l.Info("decreased reaction count")
}

func (d *Discord) onAllReactionsRemove(s *discordgo.Session, m *discordgo.MessageReactionRemoveAll) {
	// TODO: implement this
	// it seems impossible to handle this correctly, since discord.MessageReactionRemoveAll
	// does not include the number of reactions removed
}
