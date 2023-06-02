package icy

import "github.com/bwmarrin/discordgo"

type Icy struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) IcyViewer {
	return &Icy{
		ses: ses,
	}
}
