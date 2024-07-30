package tono

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TonoAdapter interface {
	GetCommunityTransaction() (*model.ListGuildCommunityTransactionResponse, error)
}
