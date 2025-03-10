package memo

import (
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Memo struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) MemoServicer {
	return &Memo{
		adapter: adapter,
		l:       l,
	}
}

func (e *Memo) GetMemos() ([]*model.Memo, error) {
	// get response from fortress
	adapterMemos, err := e.adapter.Fortress().GetMemos()
	if err != nil {
		e.l.Error(err, "can't get open memo from fortress")
		return nil, err
	}

	// normalized into in-app model
	memos := adapterMemos.Data

	return memos, nil
}

func (e *Memo) SyncMemoLogs() ([]model.MemoLog, error) {
	// sync memos from fortress
	memoLogs, err := e.adapter.Fortress().SyncMemoLogs()
	if err != nil {
		e.l.Error(err, "can't sync memo logs")
		return nil, err
	}

	// normalized into in-app model

	return memoLogs.Data, nil

}

func (e *Memo) GetMemoLogs(from, to *time.Time) ([]model.MemoLog, error) {
	// get response from fortress
	adapterMemoLogs, err := e.adapter.Fortress().GetMemoLogs(from, to)
	if err != nil {
		e.l.Error(err, "can't get memo logs from fortress")
		return nil, err
	}

	return adapterMemoLogs.Data, nil
}

func (e *Memo) GetMemoOpenPullRequest() (*model.MemoRepoWithPullRequest, error) {
	// get response from fortress
	adapterMemos, err := e.adapter.Fortress().GetMemoOpenPullRequest()
	if err != nil {
		e.l.Error(err, "can't get open memo from fortress")
		return nil, err
	}

	// normalized into in-app model
	memos := adapterMemos.Data

	return &memos, nil
}

func (e *Memo) GetMemosByDiscordID(discordID string) (*model.MemoLogsByDiscordID, error) {
	// get response from fortress
	adapterMemoLogs, err := e.adapter.Fortress().GetMemoLogsByDiscordID(discordID)
	if err != nil {
		e.l.Error(err, "can't get memos for Discord ID from fortress")
		return nil, err
	}

	return &adapterMemoLogs.Data, nil
}

func (e *Memo) GetTopAuthors(limit, days int) ([]model.AuthorRanking, error) {
	// get response from fortress
	adapterAuthors, err := e.adapter.Fortress().GetTopAuthors(limit, days)
	if err != nil {
		e.l.Error(err, "can't get top authors from fortress")
		return nil, err
	}

	return adapterAuthors.Data, nil
}
