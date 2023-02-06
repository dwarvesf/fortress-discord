package digest

import "github.com/dwarvesf/fortress-discord/pkg/model"

type DigestServicer interface {
	GetInteralUpdates() ([]*model.Digest, error)
	GetExternalUpdates() ([]*model.Digest, error)
}
