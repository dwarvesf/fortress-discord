package staff

import "github.com/bwmarrin/discordgo"

type Staff struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) StaffViewer {
	return &Staff{
		ses: ses,
	}
}
