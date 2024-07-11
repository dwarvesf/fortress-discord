package earn

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Earn) List(original *model.DiscordMessage, earns []*model.Earn) error {
	var content string
	for i := range earns {
		earn := earns[i]
		content += fmt.Sprintf("[[%d](%s)] %s\n", earn.Reward, buildEarnUrl(earn.Id), earn.Name)
	}

	msg := &discordgo.MessageEmbed{
		Title:       ":ice_cube: **Dwarves Community Earn** :ice_cube: ",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func buildEarnUrl(id string) string {
	// notion use uuid without dash
	pageId := strings.Replace(id, "-", "", -1)
	return fmt.Sprintf("https://www.notion.so/dwarves/9a5ca08b3312492b9a56cea06431842a?v=f733ece66d81446db452fc4101bdc69d&p=%s", pageId)
}

func (e *Earn) ListMemoEarn(original *model.DiscordMessage, earns []model.MemoEarn) error {
	openingContent := "<a:money:1049621199468105758> **Open**\n"
	lenOpening := 0
	doingContent := "<a:brrr:902558248907980871> **Doing**\n"
	lenDoing := 0
	for i, earn := range earns {
		content := fmt.Sprintf("[[%d](%s)] %s", i+1, earn.URL, earn.Title)

		if earn.Bounty != "" {
			content += fmt.Sprintf(" - %s", formatBounty(earn.Bounty))
		}

		var pic string
		for i := range earn.PICs {
			pic += fmt.Sprintf("<@%s> ", earn.PICs[i])
		}
		if pic != "" {
			content += fmt.Sprintf(" - %s", pic)
		}

		content += "\n"

		if strings.EqualFold(earn.Status, "Open") {
			if lenOpening >= 10 {
				continue
			}
			openingContent += content
			lenOpening++
		}

		if strings.EqualFold(earn.Status, "Doing") {
			if lenDoing >= 10 {
				continue
			}
			doingContent += content
			lenDoing++
		}
	}

	if lenDoing > 10 {
		doingContent += "\n[See more here..](https://memo.d.foundation/earn)"
	}

	if lenOpening > 10 {
		openingContent += "\n[See more here..](https://memo.d.foundation/earn)"
	}

	content := make([]string, 0)
	if lenDoing > 0 {
		content = append(content, doingContent)
	}

	if lenOpening > 0 {
		content = append(content, openingContent)
	}

	if len(content) == 0 {
		content = append(content, "There are no earning opportunities available at the moment. Please check back later!")
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:anxinicy:1014483263705862174> Bounty list <:anxinicy:1014483263705862174>",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func formatBounty(bounty string) string {
	if bounty == "" {
		return ""
	}
	return fmt.Sprintf("<a:icy:1192768878183465062> **%s ICY**", bounty)
}
