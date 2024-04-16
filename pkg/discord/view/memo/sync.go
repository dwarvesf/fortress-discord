package memo

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Memo) Sync(original *model.DiscordMessage, memos []model.MemoLog) error {
	var content string
	for i := range memos {
		if i <= 10 {
			content += fmt.Sprintf("%d ・ [%s](%s)\n", i+1, memos[i].Title, memos[i].URL)
		}
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeyes:885513213431648266> New Memo <:pepeyes:885513213431648266> \n",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
