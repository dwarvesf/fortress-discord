package discord

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

// TODO: generics this to specific packages
// will do if we have 1 more interaction, right now only support send
func (d *Discord) onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionModalSubmit {
		switch i.ModalSubmitData().CustomID {
		case "enter_advance_salary_amount_" + i.Interaction.User.ID:
			userInput := i.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

			go func() {
				salaryAdvance, err := d.Command.S.Salary().SalaryAdvance(i.Interaction.User.ID, userInput)
				if err != nil {
					d.L.Error(err, "can't make advance salary for user "+i.Interaction.User.ID)
					d.Command.View.Salary().ErrorAdvanceSalary(&model.DiscordMessage{
						ChannelId: i.ChannelID,
						Author:    i.Interaction.User,
					}, err)
					return
				}

				err = d.Command.View.Salary().CompleteAdvanceSalary(&model.DiscordMessage{
					ChannelId: i.ChannelID,
					Author:    i.Interaction.User,
				}, *salaryAdvance)
				if err != nil {
					d.L.Error(err, "can't send complete message ")
					return
				}
			}()

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						base.Normalize(s, &discordgo.MessageEmbed{
							Title: "Request Approved!\n",
							Description: fmt.Sprint(
								"Your ICY is on the way, we will notify you shortly\n\n",
								fmt.Sprintf("`Amount.  ` %s\n", fmt.Sprintf("<:ICY:1049620715374133288> **%s ICY**", userInput)),
								fmt.Sprintf("`Receiver.` %s\n", fmt.Sprintf("<@%s>", i.Interaction.User.ID)),
								"try $bals in Mochi app to see your balance",
							),
						}),
					},
				},
			})

			return
		}
	}

	if i.Type == discordgo.InteractionMessageComponent {
		var previewMode bool = !strings.Contains(i.MessageComponentData().CustomID, "--no-preview")

		// check advance salary confirm or abort button
		switch i.MessageComponentData().CustomID {
		case "open_advance_salary_button":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseModal,
				Data: &discordgo.InteractionResponseData{
					CustomID: "enter_advance_salary_amount_" + i.Interaction.User.ID,
					Title:    "Enter amount icy you want to advance",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:    "icy_amount",
									Label:       "Icy Amount",
									Style:       discordgo.TextInputShort,
									Placeholder: "100",
									Required:    true,
									MinLength:   1,
									MaxLength:   10,
								},
							},
						},
					},
				},
			})
			return
		}

		// check update type, check for "updates--" string in id
		if !strings.Contains(i.MessageComponentData().CustomID, "updates--") {
			return
		}

		// check author
		if !strings.Contains(i.MessageComponentData().CustomID, i.Member.User.ID) {
			return
		}

		// check timestamp less than 3 minutes, ignore
		if time.Since(i.Message.Timestamp).Minutes() > 3 {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: "Message too old, please try ?updates send again",
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content: "Sending updates to audiences email",
			},
		})

		// call api
		// 	curl --request POST \
		// --url 'https://develop-api.fortress.d.foundation/api/v1/notion/bc072472af5e4ab29a8025dc71565677/send?preview=true'

		// TODO:clean up
		url := "https://api.fortress.d.foundation/api/v1/notion/df-updates/%s/send"
		if previewMode {
			url = url + "?preview=true"
		}
		url = fmt.Sprintf(url, strings.Replace(i.MessageComponentData().Values[0], "-", "", -1))

		req, _ := http.NewRequest("POST", url, nil)
		req.Header.Set("Authorization", "ApiKey "+d.Cfg.ApiServer.APIKey)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		client.Do(req)
	}
}
