package hiring

import "github.com/bwmarrin/discordgo"

type Hiring struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) HiringViewer {
	return &Hiring{
		ses: ses,
	}
}
