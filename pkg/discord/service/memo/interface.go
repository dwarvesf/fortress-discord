package memo

import "github.com/dwarvesf/fortress-discord/pkg/model"

type MemoServicer interface {
	GetMemos() ([]*model.Memo, error)
}
