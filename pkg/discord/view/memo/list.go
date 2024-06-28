package memo

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (v *Memo) List(original *model.DiscordMessage, memos []*model.Memo) error {
	var content string
	for i := range memos {
		if i <= 10 {
			content += fmt.Sprintf("%d ・ [%s](%s)\n", i+1, memos[i].Name, "https://memo.d.foundation/"+strings.Replace(memos[i].Id, "-", "", -1))
		}
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepeyes:885513213431648266> Latest Memos <:pepeyes:885513213431648266> \n",
		Description: content,
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}

func (v *Memo) ListMemoLogs(original *model.DiscordMessage, memos []model.MemoLog, timeAmount int, timeUnit string) error {
	var content string

	memosByCategory := make(map[string][]model.MemoLog)
	for _, mem := range memos {
		if len(mem.Category) > 0 {
			memosByCategory[mem.Category[len(mem.Category)-1]] = append(memosByCategory[mem.Category[len(mem.Category)-1]], mem)
			continue
		}

		memosByCategory["others"] = append(memosByCategory["others"], mem)
	}

	for category, memos := range memosByCategory {
		content += fmt.Sprintf("🔹 **%s** - %v posts\n", strings.ToUpper(category), len(memos))

		tooLarge := false
		if len(memos) > 20 {
			tooLarge = true
			memos = memos[:20]
		}

		for i, memo := range memos {
			authors := make([]string, 0, len(memo.Authors))
			for _, author := range memo.Authors {
				authors = append(authors, fmt.Sprintf("<@%s>", author.DiscordID))
			}

			authorsStr := "<@anonymous>"
			if len(authors) > 0 {
				authorsStr = strings.Join(authors, ", ")
			}

			content += fmt.Sprintf("[[%d](%s)] %s - %s\n", i+1, memo.URL, memo.Title, authorsStr)
		}

		if tooLarge {
			content += "... and more"
		}

		content += "\n"
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:pepeyes:885513213431648266> Last %d %s Memos <:pepeyes:885513213431648266> \n", timeAmount, timeUnit),
		Description: content,
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
