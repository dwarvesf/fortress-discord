package memo

import (
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type MemoServicer interface {
	GetMemos() ([]*model.Memo, error)
	SyncMemoLogs() ([]model.MemoLog, error)
	GetMemoLogs(from, to *time.Time) ([]model.MemoLog, error)
	GetMemoOpenPullRequest() (*model.MemoRepoWithPullRequest, error)
	GetMemosByDiscordID(discordID string) (*model.MemoLogsByDiscordID, error)
}
