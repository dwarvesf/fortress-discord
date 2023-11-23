package salary

import (
	"github.com/bwmarrin/discordgo"
)

type Salary struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &Salary{
		ses: ses,
	}
}
