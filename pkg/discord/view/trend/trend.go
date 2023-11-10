package trend

import (
	"github.com/bwmarrin/discordgo"
)

type Trend struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) TrendViewer {
	return &Trend{
		ses: ses,
	}
}
