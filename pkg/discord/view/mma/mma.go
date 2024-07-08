package mma

import (
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MMA struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &MMA{
		ses: ses,
	}
}

func (v *MMA) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?mma template**・export csv template",
		"*Example:* `?mma template`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}

func (v *MMA) ExportTemplate(original *model.DiscordMessage, employeeMMAScores []model.EmployeeMMAScore) error {
	var csvData bytes.Buffer

	writer := csv.NewWriter(&csvData)

	// Write the header row
	header := []string{"full_name", "employee_id", "mastery_score", "autonomy_score", "meaning_score"}
	writer.Write(header)

	// Write data rows
	for _, record := range employeeMMAScores {
		data := []string{
			record.FullName,
			record.EmployeeID,
			record.MasteryScore.String(),
			record.AutonomyScore.String(),
			record.MeaningScore.String(),
		}
		writer.Write(data)
	}

	writer.Flush()

	_, err := v.ses.ChannelMessageSendComplex(original.ChannelId, &discordgo.MessageSend{
		Content: "📝 Here is the MMA template.",
		Files: []*discordgo.File{
			{
				Name:   "mma-template.csv",
				Reader: &csvData,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
