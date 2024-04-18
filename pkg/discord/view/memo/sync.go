package memo

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *Memo) Sync(original *model.DiscordMessage, memos []model.MemoLog, channelID, reward string) error {
	for i, content := range memos {
		if i <= 10 {
			var textMessage string

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

				authorField += fmt.Sprintf(" <@%s> ", author.DiscordID)
			}

			author := ""
			if authorField != "" {
				author = "from" + authorField
			}

			textMessage += fmt.Sprintf("New memo post %v \n [%s](%s)\n", author, content.Title, content.URL)

			msg := &discordgo.Message{
				Content: textMessage,
			}

			err := base.SendMessage(v.ses, original, msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
