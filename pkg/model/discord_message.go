package model

import "github.com/bwmarrin/discordgo"

type DiscordMessage struct {
	RawContent  string
	ContentArgs []string
	MessageId   string
	ChannelId   string
	GuildId     string
	Author      *discordgo.User
	Roles       []string
}
