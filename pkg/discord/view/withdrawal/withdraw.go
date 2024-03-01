package withdrawal

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Withdraw) Home(original *model.DiscordMessage, in *model.WithdrawInput, banks []model.Bank) error {
	channel, err := e.ses.UserChannelCreate(original.Author.ID)
	if err != nil {
		return err
	}

	bankOptions := make([]discordgo.SelectMenuOption, len(banks))
	for i, bank := range banks {
		bankOptions[i] = discordgo.SelectMenuOption{
			Label: bank.ShortName,
			Value: fmt.Sprintf("%v_%v", bank.SwiftCode, bank.ShortName),
			// Emoji: discordgo.ComponentEmoji{
			// 	Name: "🏦",
			// },
		}
	}

	bankSelector := discordgo.SelectMenu{
		CustomID:    "bank_selector",
		Placeholder: "Select a bank",
		Options:     bankOptions,
	}
	// if channel is direct message
	if original.ChannelId == channel.ID {
		message := discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				base.Normalize(e.ses, &discordgo.MessageEmbed{
					Title:       "ICY Withdraw",
					Description: fmt.Sprintf("You can withdraw a maximum of %s %v ICY with rate **(%vVND/ICY)**\nEnter the ICY amount :point_down:", constant.GetEmoji("ICY"), in.ICYAmount, in.ICYVNDRate),
				}),
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						bankSelector,
					},
				},
				//discordgo.ActionsRow{
				//	Components: []discordgo.MessageComponent{
				//		discordgo.Button{
				//			Label:    "Withdraw",
				//			Style:    discordgo.SecondaryButton,
				//			CustomID: "open_withdraw_form_btn",
				//			Emoji: discordgo.ComponentEmoji{
				//				Name:     "arrow_up_animated",
				//				ID:       "1131317348670902292",
				//				Animated: false,
				//			},
				//		},
				//	},
				//},
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
		Description: "?withdraw is a DM-only command. Please slide into our DM and try again.",
	}

	return base.SendEmbededMessageWithChannel(e.ses, original, msg, original.ChannelId)
}

func (s *Withdraw) ErrorWithdraw(original *model.DiscordMessage, err error) error {
	msg := &discordgo.MessageEmbed{
		Title:       "Withdraw Error!",
		Description: err.Error(),
	}

	return base.SendEmbededMessage(s.ses, original, msg)
}
