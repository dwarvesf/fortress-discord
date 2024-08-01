package tono

import "github.com/dwarvesf/fortress-discord/pkg/model"

type TonoServicer interface {
	GetCommunityTransaction() (*model.ListGuildCommunityTransaction, error)
}
