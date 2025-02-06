package techradar

import (
	"github.com/bwmarrin/discordgo"
)

type TechRadar struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) TechRadarViewer {
	return &TechRadar{ses: ses}
}
