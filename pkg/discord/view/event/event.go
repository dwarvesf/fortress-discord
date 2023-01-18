package event

import "github.com/bwmarrin/discordgo"

type Event struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) EventViewer {
	return &Event{
		ses: ses,
	}
}
