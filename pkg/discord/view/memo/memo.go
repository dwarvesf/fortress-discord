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

func (v *Memo) Help(message *model.DiscordMessage) error {
	content := []string{
		"**Memo Commands**",
		"",
		"`?memo list [duration]` - Show memo logs for a specific duration (e.g. 7d, 2w, 1m)",
		"`?memo sync` - Sync memo logs",
		"`?memo pr` - Show open pull requests",
		"`?memo top [duration]` - Show top memo authors for a specific duration",
		"  - Default: last 90 days",
		"  - Duration format examples:",
		"    • `7d` or `7 days` - Last 7 days",
		"    • `2w` or `2 weeks` - Last 2 weeks",
		"    • `1m` or `1 month` - Last month",
		"`?memo @user` - Show memos by a specific user",
		"`?memo help` - Show this help message",
		"",
		"**Aliases**:",
		"`?memos` - Alternative for memo",
		"`?memo ls` - Shorthand for list",
		"`?memo h` - Shorthand for help",
		"",
		"**Example**:",
		"`?memo list 7d` - View memo logs for the last 7 days",
		"`?memo top 1m` - View top memo authors for the last month",
		"`?memo @hnh` - View memos by a specific user",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord Bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
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
		// TODO: limit the number of posts
		if i > 4 {
			break
		}

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
		Title:       "<:pepeyes:885513213431648266> Memo Stats <:pepeyes:885513213431648266> \n",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}

func (v *Memo) ListTopAuthors(original *model.DiscordMessage, data []model.AuthorRanking, n, days int) error {
	var content []string
	for i, author := range data {
		if i >= n {
			break
		}

		authorField := ""
		if author.DiscordID != "" {
			authorField += fmt.Sprintf(" <@%s> ", author.DiscordID)
		} else if author.DiscordUsername != "" {
			authorField += fmt.Sprintf(" @%s ", author.DiscordUsername)
		} else {
			authorField += " **@unknown-user**"
		}

		content = append(content, fmt.Sprintf("%d. %v (%d memos)", i+1, authorField, author.TotalMemos))
	}
	if len(content) == 0 {
		content = append(content, "No authors found for the specified duration.")
	}
	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Top %d Memo Authors (Last %d Days)", n, days),
		Description: strings.Join(content, "\n"),
		Color:       0xED4245, // Discord red color for the left border
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "?help to see all commands",
			IconURL: "https://cdn.discordapp.com/emojis/885513213431648266.png",
		},
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
