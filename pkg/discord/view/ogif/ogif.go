package ogif

import "github.com/bwmarrin/discordgo"

type Ogif struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) OgifViewer {
	return &Ogif{
		ses: ses,
	}
}
