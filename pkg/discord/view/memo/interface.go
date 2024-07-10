package memo

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MemoViewer interface {
	List(original *model.DiscordMessage, subs []*model.Memo) error
	ListMemoLogs(original *model.DiscordMessage, subs []model.MemoLog, amount int, unit string) error
	Sync(original *model.DiscordMessage, memos []model.MemoLog, channelID, reward string) error
	ListMemoOpenPullRequest(original *model.DiscordMessage, memoPr model.MemoRepoWithPullRequest) error
}
