package issue

import "github.com/bwmarrin/discordgo"

type Issue struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) IssueViewer {
	return &Issue{
		ses: ses,
	}
}
