package df

import (
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type view struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) DFViewer {
	return &view{
		ses: ses,
	}
}

func (v *view) SendResponse(message *model.DiscordMessage, response *model.N8NEmbedResponse) error {
	fields := make([]*discordgo.MessageEmbedField, 0)
	if len(response.Fields) > 0 {
		for _, field := range response.Fields {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   field.Name,
				Value:  field.Value,
				Inline: field.Inline,
			})
		}
	}
	footer := &discordgo.MessageEmbedFooter{
		Text: response.Footer.Text,
	}

	colorValue, err := strconv.ParseInt(response.Color, 0, 32)
	if err != nil {
		return err
	}

	msg := &discordgo.MessageEmbed{
		Title:       response.Title,
		Description: response.Description,
		Color:       int(colorValue),
		Fields:      fields,
		Footer:      footer,
	}
	return base.SendEmbededMessage(v.ses, message, msg)
}
