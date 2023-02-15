package issue

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Issue struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) IssueServicer {
	return &Issue{
		adapter: adapter,
		l:       l,
	}
}

func (t *Issue) GetActiveList() ([]*model.Issue, error) {
	// send request to mochi
	adapterIssues, err := t.adapter.Fortress().GetActiveIssues()
	if err != nil {
		t.l.Error(err, "can't get issue")
		return nil, err
	}

	issues := adapterIssues.Data

	return issues, nil
}
