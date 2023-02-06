package project

import "github.com/bwmarrin/discordgo"

type Project struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) ProjectViewer {
	return &Project{
		ses: ses,
	}
}
