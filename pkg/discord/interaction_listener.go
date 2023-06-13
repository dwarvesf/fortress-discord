package discord

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// TODO: generics this to specific packages
// will do if we have 1 more interaction, right now only support send
func (d *Discord) onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionMessageComponent {
		var previewMode bool = !strings.Contains(i.MessageComponentData().CustomID, "--no-preview")

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
