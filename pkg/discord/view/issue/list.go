package issue

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (i *Issue) ListActive(original *model.DiscordMessage, issues []*model.Issue) error {
	var content string

	for i := range issues {
		content += fmt.Sprintf("%d・ %s\n", i+1, issues[i].Name)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeno2:885513214467661834> Current Issues <:pepeno2:885513214467661834>",
		Description: content,
	}

	return base.SendEmbededMessage(i.ses, original, msg)
}
