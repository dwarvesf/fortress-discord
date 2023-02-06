package digest

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DigestViewer interface {
	ListInternal(original *model.DiscordMessage, digest []*model.Digest) error
	ListExternal(original *model.DiscordMessage, digest []*model.Digest) error
}
