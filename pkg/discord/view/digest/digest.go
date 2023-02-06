package digest

import "github.com/bwmarrin/discordgo"

type Digest struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) DigestViewer {
	return &Digest{
		ses: ses,
	}
}
