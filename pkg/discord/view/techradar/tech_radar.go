package techradar

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type TechRadar struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) TechRadarViewer {
	return &TechRadar{ses: ses}
}

func (t *TechRadar) Help(original *model.DiscordMessage) error {
	helpMessage := "**Tech Radar Commands:**\n" +
		"- `/tech-radar list trial`: List technologies in the Trial quadrant\n" +
		"- `/tech-radar list assess`: List technologies in the Assess quadrant\n" +
		"- `/tech-radar list adopt`: List technologies in the Adopt quadrant\n" +
		"- `/tech-radar list hold`: List technologies in the Hold quadrant\n" +
		"- `/tech-radar search <keyword>`: Search for technologies by keyword\n\n" +
		"**Usage Tips:**\n" +
		"- Use these commands to explore and track technology trends\n" +
		"- Each quadrant represents a different stage of technology adoption\n"

	_, err := t.ses.ChannelMessageSend(original.ChannelId, helpMessage)
	return err
}
