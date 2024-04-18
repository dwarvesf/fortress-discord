package memo

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *Memo) Sync(original *model.DiscordMessage, memos []model.MemoLog, channelID, reward string) error {
	for _, content := range memos {
		authorField := ""
		avatar := ""
		for _, author := range content.Authors {
			if avatar == "" {
				discordMember, err := v.ses.GuildMember(original.GuildId, content.Authors[0].DiscordID)
				if err != nil {
					return err
				}

				avatar = fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.webp?size=240", discordMember.User.ID, discordMember.User.Avatar)
				if discordMember.Avatar != "" {
					avatar = fmt.Sprintf("https://cdn.discordapp.com/guilds/%v/users/%v/avatars/%v.webp?size=240", original.GuildId, discordMember.User.ID, discordMember.Avatar)
				}
			}

			authorField += fmt.Sprintf(" <@%s>\n", author.DiscordID)
		}

		if authorField == "" {
			authorField = "N/A"
		}

		messageEmbed := []*discordgo.MessageEmbedField{
			{
				Name:   "Author",
				Value:  authorField,
				Inline: true,
			},
			{
				Name:   "ICY 🧊",
				Value:  reward,
				Inline: true,
			},
		}

		msg := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s", content.Title),
			URL:         content.URL,
			Description: content.Description,
			Fields:      messageEmbed,
			Timestamp:   content.PublishedAt.Format(time.RFC3339),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: avatar,
			},
		}
		err := base.SendEmbededMessageWithChannel(v.ses, original, msg, channelID)
		if err != nil {
			return err
		}
	}

	return nil
}
