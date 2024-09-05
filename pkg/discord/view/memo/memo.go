package memo

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Memo struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) MemoViewer {
	return &Memo{
		ses: ses,
	}
}

func (v *Memo) ListByDiscordID(original *model.DiscordMessage, data *model.MemoLogsByDiscordID, discordID string) error {
	content := []string{
		fmt.Sprintf("**Memos created by <@%s>**", discordID),
		fmt.Sprintf("`Total:  `%v posts", data.Rank.TotalMemos),
		fmt.Sprintf("`Rank:   `#%v", data.Rank.Rank),
		"",
	}

	postStr := "**Latest Memos**\n"
	for i, memo := range data.MemoLogs {
		authors := make([]string, 0, len(memo.Authors))
		for _, author := range memo.Authors {
			authors = append(authors, fmt.Sprintf("<@%s>", author.DiscordID))
		}

		authorsStr := "**@unknown-user**"
		if len(authors) > 0 {
			authorsStr = strings.Join(authors, ", ")
		}

		postStr += fmt.Sprintf("[[%d](%s)] %s - %s\n", i+1, memo.URL, memo.Title, authorsStr)
	}

	if postStr != "" {

		content = append(content, postStr)
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:pepeyes:885513213431648266> Memo Stats <:pepeyes:885513213431648266> \n"),
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}

func (v *Memo) ListTopAuthors(original *model.DiscordMessage, data []model.AuthorRanking) error {
	var content []string

	for i, author := range data {
		content = append(content, fmt.Sprintf("[[%v]](%s) <@%s> - %v posts", i+1, fmt.Sprintf("https://memo.d.foundation/contributor/%s", author.MemoUsername), author.DiscordID, author.TotalMemos))
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeyes:885513213431648266> Memo Leaderboard <:pepeyes:885513213431648266> \n",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
