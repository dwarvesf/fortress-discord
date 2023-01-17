package errors

import "github.com/bwmarrin/discordgo"

type Error struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) ErrorViewer {
	return &Error{
		ses: ses,
	}
}
