package memo

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

var (
	memoCategoryList = []string{memoCategoryFleeting, memoCategoryLiterature, memoCategoryEarn, memoCategoryOthers}
)

const (
	memoCategoryFleeting   = "00_fleeting"
	memoCategoryLiterature = "01_literature"
	memoCategoryEarn       = "earn"
	memoCategoryOthers     = "others"
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

	//Group by category
	memosByCategory := map[string][]model.MemoLog{
		memoCategoryFleeting:   make([]model.MemoLog, 0),
		memoCategoryLiterature: make([]model.MemoLog, 0),
		memoCategoryEarn:       make([]model.MemoLog, 0),
		memoCategoryOthers:     make([]model.MemoLog, 0),
	}

	for _, mem := range memos {
		isMapped := false
		for _, category := range mem.Category {
			if strings.EqualFold(category, memoCategoryFleeting) ||
				strings.EqualFold(category, memoCategoryLiterature) ||
				strings.EqualFold(category, memoCategoryEarn) {
				memosByCategory[category] = append(memosByCategory[category], mem)
				isMapped = true
				break
			}
		}

		if !isMapped {
			memosByCategory[memoCategoryOthers] = append(memosByCategory[memoCategoryOthers], mem)
		}
	}

	for _, category := range memoCategoryList {
		memos := memosByCategory[category]
		content += fmt.Sprintf("🔹 **%s** - %v posts\n", strings.ToUpper(category), len(memos))

		tooLarge := false
		if len(memos) > 7 {
			tooLarge = true
			memos = memos[:7]
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
			content += "... more\n"
		}

		content += "\n"
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:pepeyes:885513213431648266> Last %d %s Memos <:pepeyes:885513213431648266> \n", timeAmount, timeUnit),
		Description: content,
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
