package digest

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Digest) ListInternal(original *model.DiscordMessage, digests []*model.Digest) error {
	var content string
	for i := range digests {
		s := strings.Split(digests[i].Name, "-")
		if len(s) == 2 {
			content += fmt.Sprintf("[%s](%s) ・ %s \n", strings.TrimSpace(s[0]), fmt.Sprintf("https://digest.d.foundation/%s", strings.ReplaceAll(digests[i].Id, "-", "")), strings.TrimSpace(s[1]))
		}
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeyes:885513213431648266> Latest Digests <:pepeyes:885513213431648266> \n",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func (e *Digest) ListExternal(original *model.DiscordMessage, digests []*model.Digest) error {
	var content string
	for i := range digests {
		content += fmt.Sprintf("%d ・[%s](%s)\n", i+1, digests[i].Name, fmt.Sprintf("https://digest.d.foundation/%s", strings.ReplaceAll(digests[i].Id, "-", "")))
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeyes:885513213431648266> Latest Team Updates <:pepeyes:885513213431648266> \n",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}
