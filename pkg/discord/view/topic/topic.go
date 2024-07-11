package topic

import "github.com/bwmarrin/discordgo"

type Topic struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) TopicViewer {
	return &Topic{
		ses: ses,
	}
}
