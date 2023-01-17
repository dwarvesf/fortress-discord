package earn

import "github.com/bwmarrin/discordgo"

type Earn struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) EarnViewer {
	return &Earn{
		ses: ses,
	}
}
