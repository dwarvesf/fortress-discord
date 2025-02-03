package salary

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Salary struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &Salary{
		ses: ses,
	}
}

func (s *Salary) Help(original *model.DiscordMessage) error {
	helpMessage := "**Salary Commands:**\n" +
		"- `/salary advance`: Advance your salary\n" +
		"- `/salary help`: Show this message\n\n" +
		"**Usage Tips:**\n" +
		"- Use these commands to manage your salary\n"

	_, err := s.ses.ChannelMessageSend(original.ChannelId, helpMessage)
	return err
}
