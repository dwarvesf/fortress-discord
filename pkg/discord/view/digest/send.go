package digest

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (d *Digest) SendoutSelection(original *model.DiscordMessage, digest []*model.Digest) error {

	var previewMode bool = !strings.Contains(original.RawContent, "--no-preview")

	opts := []discordgo.SelectMenuOption{}
	for i := range digest {
		opts = append(opts, discordgo.SelectMenuOption{
			Label: digest[i].Name,
			Value: digest[i].Id,
		})
	}

	// id = updates--<user_id>--<timestamp>[--<no-preview>]
	id := fmt.Sprintf("updates--%s--%d", original.Author.ID, time.Now().Unix())
	if !previewMode {
		id += "--no-preview"
	}

	selectMenu := &discordgo.SelectMenu{
		CustomID:    id,
		Options:     opts,
		Placeholder: "Select an updates to send",
	}

	content := "Select Updates to send"
	if previewMode {
		content += "\n\nwe will send this update as preview mode, only to small specifics email pre-configured. \n\n"
		content += "If you want to send this update to all audiences, please use `?updates send --no-preview`"
	}
	if !previewMode {
		content += "\n we will send this update to **ALL AUDIENCES**, please use `?updates send` to send as preview mode"
	}

	message := &discordgo.MessageSend{
		Content: content,
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					selectMenu,
				},
			},
		},
	}

	d.ses.ChannelMessageSendComplex(original.ChannelId, message)
	return nil
}
