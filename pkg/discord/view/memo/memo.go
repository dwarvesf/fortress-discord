package memo

import "github.com/bwmarrin/discordgo"

type Memo struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) MemoViewer {
	return &Memo{
		ses: ses,
	}
}
