package discord

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// onMessageCreate is an entry point for discord message create (chat) event
func (d *Discord) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Forward pull request messages from the "dev" channel to the "random" channel
	d.forwardPullRequestMessage(s, m)

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

func (d *Discord) onGuildScheduledEventCreate(s *discordgo.Session, m *discordgo.GuildScheduledEventCreate) {
	if err := d.Command.S.Event().CreateGuildScheduledEvent(&model.DiscordEvent{
		ID:               m.ID,
		DiscordChannelID: m.ChannelID,
		DiscordCreatorID: m.CreatorID,
		Name:             m.Name,
		Description:      m.Description,
		Date:             m.ScheduledStartTime,
	}); err != nil {
		d.L.Error(err, "failed to create a scheduled event on discord")
	}
}

func (d *Discord) forwardPullRequestMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if message is from the "dev" channel
	if m.ChannelID != d.Cfg.Discord.ID.DevChannel {
		return
	}

	// check message create from bot
	if m.Author != nil && !m.Author.Bot {
		return
	}

	content := getDiscordEmbedTitle(m.Message)

	// Check if message content indicates a new pull request
	if !strings.Contains(content, "Pull request opened:") {
		return
	}

	newMessage := &discordgo.MessageSend{
		Embeds: m.Embeds,
	}

	// Send the message to the "random" channel
	_, err := s.ChannelMessageSendComplex(d.Cfg.Discord.ID.RandomChannel, newMessage)
	if err != nil {
		log.Printf("Error sending message to random channel: %v", err)
	}
}

func getDiscordEmbedTitle(m *discordgo.Message) string {
	if len(m.Embeds) == 0 {
		return ""
	}

	// Assuming we're interested in the first embed
	embed := m.Embeds[0]

	// If no description, check for title
	if embed.Title != "" {
		return embed.Title
	}

	return ""
}
