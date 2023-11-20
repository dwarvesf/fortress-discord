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
				base.Normalize(original.Author, &discordgo.MessageEmbed{
					Title:       "Enter ICY amount you want to borrow",
					Description: fmt.Sprintf("The maximum ICY you can borrow: :icy: %s ICY ($%s)\nPlease enter the ICY amount :point_down:", checkSalaryAdvance.AmountIcy, checkSalaryAdvance.AmountUSD),
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
		Title:       "DM Fortress",
		Description: "?salary advance is only used in DM with Fortress! Please send a message to Fortress bot!",
	}

	return base.SendEmbededMessageWithChannel(e.ses, original, msg, channel.ID)
}

// CompleteAdvanceSalary implements Viewer.
func (s *Salary) CompleteAdvanceSalary(original *model.DiscordMessage, salaryAdvance model.SalaryAdvance) error {
	msg := &discordgo.MessageEmbed{
		Title: "<a:star_animated:1131862886592024586> Successfully Salary advance",
		Description: fmt.Sprint(
			fmt.Sprintf("`TxID.    ` %s\n", fmt.Sprintf("[%s](https://mochi.gg/tx/%s)", salaryAdvance.TransactionID, salaryAdvance.TransactionID)),
			fmt.Sprintf("`TxHash.  ` %s\n", fmt.Sprintf("[%s](https://polygonscan.com/tx/%s)", salaryAdvance.TransactionHash, salaryAdvance.TransactionHash)),
			fmt.Sprintf("`Amount.  ` %s\n", fmt.Sprintf("<:ICY:1049620715374133288> **%s icy** (%s USD)", salaryAdvance.AmountIcy, salaryAdvance.AmountUSD)),
			fmt.Sprintf("`Sender.  ` %s\n", "**Dwarves Foundation Salary Advance**"),
			fmt.Sprintf("`Receiver.` %s\n", fmt.Sprintf("<@%s>", original.Author.ID)),
		),
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}

func (s *Salary) ErrorAdvanceSalary(original *model.DiscordMessage) error {
	msg := &discordgo.MessageEmbed{
		Title:       "Ok",
		Description: "Ok",
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
