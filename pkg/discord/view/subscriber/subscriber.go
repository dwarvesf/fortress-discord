package subscriber

import "github.com/bwmarrin/discordgo"

type Subscriber struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) SubscriberViewer {
	return &Subscriber{
		ses: ses,
	}
}
