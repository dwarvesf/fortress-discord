package issue

import "github.com/dwarvesf/fortress-discord/pkg/model"

type IssueServicer interface {
	GetActiveList() ([]*model.Issue, error)
}
