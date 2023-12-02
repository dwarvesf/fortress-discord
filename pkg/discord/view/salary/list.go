package salary

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Salary) EnterAmountAdvanceSalary(original *model.DiscordMessage, checkSalaryAdvance model.CheckSalaryAdvance) error {
	channel, err := e.ses.UserChannelCreate(original.Author.ID)
	if err != nil {
		return err
	}

	// if channel is direct message
	if original.ChannelId == channel.ID {
		message := discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				base.Normalize(e.ses, &discordgo.MessageEmbed{
					Title:       "Enter amount",
					Description: fmt.Sprintf("You can advance a maximum of: <:ICY:1049620715374133288> %s ICY (%s) credit\nthe given credit will be paid automatically in your next pay check.\nEnter the ICY amount :point_down:", checkSalaryAdvance.AmountICY, checkSalaryAdvance.AmountUSD),
				}),
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Enter",
							Style:    discordgo.SecondaryButton,
							CustomID: "open_advance_salary_button",
						},
					},
				},
			},
		}

		_, err = e.ses.ChannelMessageSendComplex(channel.ID, &message)
		if err != nil {
			return err
		}
		return nil
	}

	// if public channel
	msg := &discordgo.MessageEmbed{
		Title:       "Fortress",
		Description: "?salary advance is a DM-only command. Please slide into our DM and try again.",
	}

	return base.SendEmbededMessageWithChannel(e.ses, original, msg, original.ChannelId)
}

// CompleteAdvanceSalary implements Viewer.
func (s *Salary) CompleteAdvanceSalary(original *model.DiscordMessage, salaryAdvance model.SalaryAdvance) error {
	msg := &discordgo.MessageEmbed{
		Title: "<a:star_animated:1131862886592024586> Successfully Transaction",
		Description: fmt.Sprint(
			fmt.Sprintf("`TxID.    ` %s\n", fmt.Sprintf("[%s](https://mochi.gg/tx/%s)", salaryAdvance.TransactionHash, salaryAdvance.TransactionHash)),
			fmt.Sprintf("`Amount.  ` %s\n", fmt.Sprintf("<:ICY:1049620715374133288> **%s ICY** (%s)", salaryAdvance.AmountICY, salaryAdvance.AmountUSD)),
			fmt.Sprintf("`Sender.  ` %s\n", "**Dwarves Foundation**"),
			fmt.Sprintf("`Receiver.` %s\n", fmt.Sprintf("<@%s>", original.Author.ID)),
		),
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}

func (s *Salary) ErrorAdvanceSalary(original *model.DiscordMessage, err error) error {
	msg := &discordgo.MessageEmbed{
		Title:       "Salary Advance Error!",
		Description: err.Error(),
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
