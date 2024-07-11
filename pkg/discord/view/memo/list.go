package memo

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/stringutils"
)

var (
	memoCategoryList = []string{memoCategoryFleeting, memoCategoryLiterature, memoCategoryEarn, memoCategoryOthers}
	githubUrl        = "https://github.com"
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

			authorsStr := "**@unknown-user**"
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

func (v *Memo) ListMemoOpenPullRequest(original *model.DiscordMessage, memoPr model.MemoRepoWithPullRequest) error {
	content := ""

	// get repo
	repos := stringutils.GetKeysFromMap(memoPr)
	stringutils.SortSlice(repos)

	for _, repo := range repos {
		prs := memoPr[repo]

		if len(prs) == 0 {
			continue
		}

		content += fmt.Sprintf("**%s** \n", repo)

		for _, pr := range prs {
			author := fmt.Sprintf("[%s](%s/%s)", pr.GithubUserName, githubUrl, pr.GithubUserName)
			if pr.DiscordId != "" {
				author = fmt.Sprintf("<@%s>", pr.DiscordId)
			}
			content += fmt.Sprintf("∟ %s [[#%d](%s)] %s - %s \n", stringutils.ConvertToTimeAgo(pr.Timestamp), pr.Number, pr.Url, pr.Title, author)
		}

		content += "\n"
	}

	if content == "" {
		content = "No open PRs found"
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:pepe_ping:1028964391690965012> Memos PR List <:pepe_ping:1028964391690965012>\n"),
		Description: content,
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}
