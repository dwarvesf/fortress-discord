package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MemoViewer interface {
	Help(original *model.DiscordMessage) error
	List(original *model.DiscordMessage, subs []*model.Memo) error
	ListMemoLogs(original *model.DiscordMessage, subs []model.MemoLog, amount int, unit string) error
	Sync(original *model.DiscordMessage, memos []model.MemoLog, channelID, reward string) error
	ListMemoOpenPullRequest(original *model.DiscordMessage, memoPr model.MemoRepoWithPullRequest) error
	ListByDiscordID(original *model.DiscordMessage, data *model.MemoLogsByDiscordID, discordID string) error
	ListTopAuthors(original *model.DiscordMessage, data []model.AuthorRanking, n, days int) error
}
