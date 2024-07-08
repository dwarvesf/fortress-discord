package news

import "github.com/bwmarrin/discordgo"

type view struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &view{
		ses: ses,
	}
}
